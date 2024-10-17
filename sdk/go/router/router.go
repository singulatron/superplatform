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

	"github.com/singulatron/superplatform/sdk/go/datastore"
)

var port = "58231"

const defaultAddress = "http://127.0.0.1"

func GetPort() string {
	return port
}

func SetPort(i int) {
	port = fmt.Sprintf("%v", i)
}

func SelfAddress() string {
	return fmt.Sprintf("%v:%v", defaultAddress, GetPort())
}

type Router struct {
	registry      map[string]string
	address       string
	bearerToken   string
	mockEndpoints map[string]any
}

func NewRouter(datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) (*Router, error) {
	return &Router{
		registry:      map[string]string{},
		mockEndpoints: map[string]any{},
	}, nil
}

func (r *Router) SetDefaultAddress(address string) {
	r.address = address
}

func (r *Router) Address() string {
	if r.address != "" {
		return r.address
	}

	return fmt.Sprintf("%v:%v", defaultAddress, port)
}

func (r *Router) AddMock(serviceName, path string, rsp any) {
	r.mockEndpoints[serviceName+path] = rsp
}

func (r *Router) SetBearerToken(token string) *Router {
	return &Router{
		registry:      r.registry,
		address:       r.address,
		mockEndpoints: r.mockEndpoints,
		bearerToken:   token,
	}
}

func (r *Router) AsRequestMaker(req *http.Request) *Router {
	authHeader := req.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" {
		return r.SetBearerToken("")
	}

	return r.SetBearerToken(authHeader)
}

func (r *Router) request(ctx context.Context, method, serviceName, path string, requestBody, responseBody any) error {
	val, ok := r.mockEndpoints[serviceName+path]
	if ok {
		bs, _ := json.Marshal(val)
		return json.Unmarshal(bs, responseBody)
	}

	address, ok := r.registry[serviceName]
	if !ok {
		address = r.Address()
	}

	var bodyBytes []byte
	var err error
	if requestBody != nil {
		bodyBytes, err = json.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	url := fmt.Sprintf("%v/%v%v", address, serviceName, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bodyBytes))
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

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse map[string]string
		if err := json.Unmarshal(responseBytes, &errResponse); err != nil {
			return formatError(fmt.Errorf("%s %v -> '%s': %v", method, url, resp.Status, string(responseBytes)))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return formatError(fmt.Errorf("%s %v -> '%s': %v", method, url, resp.Status, errMsg))
		}
		return formatError(fmt.Errorf("%s %v -> '%s': %v", method, url, resp.Status, string(responseBytes)))
	}

	if responseBody != nil {
		err = json.Unmarshal(responseBytes, responseBody)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Router) Post(ctx context.Context, serviceName, path string, request, response any) error {
	return r.request(ctx, "POST", serviceName, path, request, response)
}

func (r *Router) Get(ctx context.Context, serviceName, path string, queryParams map[string]string, response any) error {
	val, ok := r.mockEndpoints[serviceName+path]
	if ok {
		bs, _ := json.Marshal(val)
		return json.Unmarshal(bs, response)
	}

	address, ok := r.registry[serviceName]
	if !ok {
		address = r.Address()
	}

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

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse map[string]string
		if err := json.Unmarshal(responseBytes, &errResponse); err != nil {
			return formatError(fmt.Errorf("GET %v -> '%s': %v", ur, resp.Status, string(responseBytes)))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return formatError(fmt.Errorf("GET %v -> '%s': %v", ur, resp.Status, errMsg))
		}
		return formatError(fmt.Errorf("GET %v -> '%s': %v", ur, resp.Status, string(responseBytes)))
	}

	if response != nil {
		err = json.Unmarshal(responseBytes, response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Router) Put(ctx context.Context, serviceName, path string, request, response any) error {
	return r.request(ctx, "PUT", serviceName, path, request, response)
}

func (r *Router) Delete(ctx context.Context, serviceName, path string, request, response any) error {
	return r.request(ctx, "DELETE", serviceName, path, request, response)
}

func formatError(err error) error {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return fmt.Errorf("error: %w", err)
	}
	return fmt.Errorf("%s:%d: %w", file, line, err)
}
