package api

import (
	"fmt"
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

func (c *ApiClient) url(resource string) string {
	return fmt.Sprintf("%s%s", c.baseUrl, resource)
}

func (c *ApiClient) Do(options RequestOptions) *ApiClient {
	request, err := http.NewRequest(options.Method, c.url(options.Path), options.Body)

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
