package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) {
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
			return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
		}
		if errMsg, exists := errResponse["error"]; exists {
			return fmt.Errorf("service error: %s", errMsg)
		}
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	if response != nil {
		err = json.Unmarshal(responseBody, response)
		if err != nil {
			return err
		}
	}

	return nil
}
