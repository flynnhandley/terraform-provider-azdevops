package azdevops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Error represents a error from the visual studio api.
type Error struct {
	APIError struct {
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	Type       string `json:"type,omitempty"`
	StatusCode int
	Endpoint   string
}

func (e Error) Error() string {
	return fmt.Sprintf("API Error: %d %s %s", e.StatusCode, e.Endpoint, e.APIError.Message)
}

type AzDevOpsClient struct {
	Account    string
	Token      string
	HTTPClient *http.Client
}

const AcceptHeader = "application/json;api-version=4.0-preview;excludeUrls=true"

func (c *AzDevOpsClient) Do(method, endpoint string, payload *bytes.Buffer) (*http.Response, error) {

	absoluteendpoint := fmt.Sprintf("https://dev.azure.com/%s/%s", c.Account, endpoint)
	log.Printf("[DEBUG] Sending request to %s %s", method, absoluteendpoint)

	var bodyreader io.Reader

	if payload != nil {
		log.Printf("[DEBUG] With payload %s", payload.String())
		bodyreader = payload
	}

	req, err := http.NewRequest(method, absoluteendpoint, bodyreader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", AcceptHeader)

	req.Header.Add("Authorization", "Basic <<TODO: Add basic auth here>>")

	if payload != nil {
		// Can cause bad request when putting default reviews if set.
		req.Header.Add("Content-Type", "application/json")
	}

	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	log.Printf("[DEBUG] Resp: %v Err: %v", resp, err)
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		apiError := Error{
			StatusCode: resp.StatusCode,
			Endpoint:   endpoint,
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		log.Printf("[DEBUG] Resp Body: %s", string(body))

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			apiError.APIError.Message = string(body)
		}

		return resp, error(apiError)

	}
	return resp, err
}

func (c *AzDevOpsClient) Get(endpoint string) (*http.Response, error) {
	return c.Do("GET", endpoint, nil)
}

func (c *AzDevOpsClient) Patch(endpoint string, jsonpayload *bytes.Buffer) (*http.Response, error) {
	return c.Do("PATCH", endpoint, jsonpayload)
}

func (c *AzDevOpsClient) Post(endpoint string, jsonpayload *bytes.Buffer) (*http.Response, error) {
	return c.Do("POST", endpoint, jsonpayload)
}

func (c *AzDevOpsClient) Put(endpoint string, jsonpayload *bytes.Buffer) (*http.Response, error) {
	return c.Do("PUT", endpoint, jsonpayload)
}

func (c *AzDevOpsClient) PutOnly(endpoint string) (*http.Response, error) {
	return c.Do("PUT", endpoint, nil)
}

func (c *AzDevOpsClient) Delete(endpoint string) (*http.Response, error) {
	return c.Do("DELETE", endpoint, nil)
}
