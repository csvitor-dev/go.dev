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
	httpClient *http.Client
	token      string
	response   *http.Response
	clientErr  error
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
		httpClient: &http.Client{
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

func (c *ApiClient) Do(options RequestOptions) *ApiClient {
	var reqBody io.Reader
	jsonBody, err := json.Marshal(options.Body)

	if err != nil {
		c.clientErr = err
		return c
	}
	reqBody = bytes.NewBuffer(jsonBody)
	request, err := http.NewRequest(options.Method, c.url(options.Path), reqBody)

	if err != nil {
		c.clientErr = err
		return c
	}
	request.Header.Set("Content-Type", "application/json")

	if options.RequireAuth {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	c.response, c.clientErr = c.httpClient.Do(request)

	return c
}

func (c *ApiClient) Done() (*http.Response, error) {
	return c.response, c.clientErr
}
