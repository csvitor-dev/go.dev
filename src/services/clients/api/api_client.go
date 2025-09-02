package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiClient struct {
	baseUrl    string
	HttpClient *http.Client
	token      string
}

type RequestOptions struct {
	Body        any
	RequireAuth bool
	Method      string
	Path        string
}

func NewApiClient(baseUrl string) *ApiClient {
	return &ApiClient{
		baseUrl: baseUrl,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *ApiClient) WithToken(token string) *ApiClient {
	c.token = token
	return c
}

func (c *ApiClient) url(path string) string {
	return fmt.Sprintf("%s%s", c.baseUrl, path)
}

func ExecuteRequest[T any](client *ApiClient, options RequestOptions) (T, error) {
	var response T
	payload, err := client.do(options)

	if err != nil || len(payload) == 0 {
		return response, err
	}
	err = json.Unmarshal(payload, &response)

	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *ApiClient) do(options RequestOptions) ([]byte, error) {
	var reqBody io.Reader
	jsonBody, err := json.Marshal(options.Body)

	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewBuffer(jsonBody)
	request, err := http.NewRequest(options.Method, c.url(options.Path), reqBody)

	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	if options.RequireAuth {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	response, err := c.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("API error (%d): %s", response.StatusCode, string(data))
	}
	return data, nil
}
