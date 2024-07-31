package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/singulatron/singulatron/localtron/datastore"
)

const DefaultPort = "58231"

type Router struct {
	registry       map[string]string
	defaultAddress string
	bearerToken    string
	mockEndpoints  map[string]any
}

func NewRouter(
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) (*Router, error) {
	return &Router{
		registry:       map[string]string{},
		defaultAddress: "http://127.0.0.1:" + DefaultPort,
		mockEndpoints:  map[string]any{},
	}, nil
}

func (r *Router) SetDefaultAddress(address string) {
	r.defaultAddress = address
}

func (r *Router) AddMock(serviceName, path string, rsp any) {
	r.mockEndpoints[serviceName+path] = rsp
}

func (r *Router) SetBearerToken(token string) *Router {
	return &Router{
		// @todo copy?
		registry: r.registry,

		defaultAddress: r.defaultAddress,
		mockEndpoints:  r.mockEndpoints,
		bearerToken:    token,
	}
}

func (router *Router) AsRequestMaker(r *http.Request) *Router {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" {
		return router.SetBearerToken("")
	}

	return router.SetBearerToken(authHeader)
}

func (r *Router) Post(ctx context.Context, serviceName, path string, request, response any) error {
	val, ok := r.mockEndpoints[serviceName+path]
	if ok {
		bs, _ := json.Marshal(val)
		return json.Unmarshal(bs, response)
	}

	address, ok := r.registry[serviceName]
	if !ok {
		address = r.defaultAddress
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/%v%v", address, serviceName, path)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse map[string]string
		if err := json.Unmarshal(responseBody, &errResponse); err != nil {
			return formatError(fmt.Errorf("POST %v -> '%s': %v", url, resp.Status, string(responseBody)))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return formatError(fmt.Errorf("POST %v -> '%s': %v", url, resp.Status, errMsg))
		}
		return formatError(fmt.Errorf("POST %v -> '%s': %v", url, resp.Status, string(responseBody)))
	}

	if response != nil {
		err = json.Unmarshal(responseBody, response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Router) Get(ctx context.Context, serviceName, path string, queryParams map[string]string, response any) error {
	val, ok := r.mockEndpoints[serviceName+path]
	if ok {
		bs, _ := json.Marshal(val)
		return json.Unmarshal(bs, response)
	}

	address, ok := r.registry[serviceName]
	if !ok {
		address = r.defaultAddress
	}

	// Construct URL with query parameters
	ur := fmt.Sprintf("%v/%v%v", address, serviceName, path)
	if len(queryParams) > 0 {
		q := url.Values{}
		for key, value := range queryParams {
			q.Add(key, value)
		}
		ur += "?" + q.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", ur, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse map[string]string
		if err := json.Unmarshal(responseBody, &errResponse); err != nil {
			return formatError(fmt.Errorf("GET %v -> '%s': %v", ur, resp.Status, string(responseBody)))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return formatError(fmt.Errorf("GET %v -> '%s': %v", ur, resp.Status, errMsg))
		}
		return formatError(fmt.Errorf("GET %v ->'%s': %v", ur, resp.Status, string(responseBody)))
	}

	if response != nil {
		err = json.Unmarshal(responseBody, response)
		if err != nil {
			return err
		}
	}

	return nil
}

func formatError(err error) error {
	return err

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return fmt.Errorf("error: %w", err)
	}
	return fmt.Errorf("%s:%d: %w", file, line, err)
}
