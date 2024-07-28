package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/singulatron/singulatron/localtron/datastore"
)

const DefaultPort = "58231"

type RouteContext struct {
	Context context.Context
	Headers http.Header
}

type Router struct {
	registry map[string]string
}

func NewRouter(
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) (*Router, error) {
	return &Router{
		registry: map[string]string{},
	}, nil
}

func (r *Router) Post(ctx context.Context, serviceName, path string, request, response any) error {
	address, ok := r.registry[serviceName]
	if !ok {
		address = "http://127.0.0.1:" + DefaultPort
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v%v", address, path)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

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
			return fmt.Errorf("service '%s' %s returned non-OK HTTP status: %s, body: %v", serviceName, path, resp.Status, string(responseBody))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return fmt.Errorf("service '%s' %s error: %s", serviceName, path, errMsg)
		}
		return fmt.Errorf("service '%s' %s returned non-OK HTTP status: %s, body: %v", serviceName, path, resp.Status, string(responseBody))
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
	address, ok := r.registry[serviceName]
	if !ok {
		address = "http://127.0.0.1:" + DefaultPort
	}

	// Construct URL with query parameters
	ur := fmt.Sprintf("%v%v", address, path)
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
			return fmt.Errorf("service '%s' %s returned non-OK HTTP status: %s, body: %v", serviceName, path, resp.Status, string(responseBody))
		}
		if errMsg, exists := errResponse["error"]; exists {
			return fmt.Errorf("service '%s' %s error: %s", serviceName, path, errMsg)
		}
		return fmt.Errorf("service '%s' %s returned non-OK HTTP status: %s, body: %v", serviceName, path, resp.Status, string(responseBody))
	}

	if response != nil {
		err = json.Unmarshal(responseBody, response)
		if err != nil {
			return err
		}
	}

	return nil
}
