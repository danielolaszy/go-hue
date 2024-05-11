package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hue/utils"
	"net/http"
)

// Client represents the API client for interacting with the Philips Hue API.
type Client struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

// NewClient creates a new instance of the API client.
func NewClient(baseURL, authToken string) *Client {
	return &Client{
		baseURL:    baseURL,
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

// makeRequest makes an HTTP request to the Philips Hue API.
func (c *Client) MakeRequest(method, path string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	// Construct URL
	url := utils.BuildURL(c.baseURL, path, queryParams)

	// Marshal request body to JSON
	var requestBody []byte
	if body != nil {
		requestBody, _ = json.Marshal(body)
	}

	// Create HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Set authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	// Send HTTP request
	return c.httpClient.Do(req)
}
