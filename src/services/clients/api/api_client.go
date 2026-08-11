package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	res "github.com/csvitor-dev/go.dev/pkg/responses"
)

type ApiClient struct {
	BaseUrl    string
	Writer     http.ResponseWriter
	HttpClient *http.Client
	Token      string
}

func NewApiClient(baseUrl string, writer http.ResponseWriter, token ...string) *ApiClient {
	var onlyToken string

	if len(token) == 1 {
		onlyToken = token[0]
	}
	return &ApiClient{
		BaseUrl: baseUrl,
		Writer:  writer,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Token: onlyToken,
	}
}

func (c *ApiClient) url(path string) string {
	return fmt.Sprintf("%s%s", c.BaseUrl, path)
}

func ExecuteRequest[T any](client *ApiClient, options RequestOptions) (T, int) {
	var response T
	payload, status := client.do(options)

	if !parse(client.Writer, payload, &response) {
		return response, 0
	}
	return response, status
}

func (c *ApiClient) do(options RequestOptions) ([]byte, int) {
	var reqBody io.Reader
	jsonBody, err := json.Marshal(options.Body)

	if err != nil {
		res.SingleError(c.Writer, 0, errors.New("Falha ao gerar parse da requisição."))
		return nil, 0
	}
	reqBody = bytes.NewBuffer(jsonBody)
	request, err := http.NewRequest(options.Method, c.url(options.Path), reqBody)

	if err != nil {
		res.SingleError(c.Writer, 0, errors.New("Falha na tentativa de criar a requisição."))
		return nil, 0
	}
	request.Header.Set("Content-Type", "application/json")

	if options.IsAuthenticated {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	response, err := c.HttpClient.Do(request)

	if err != nil {
		res.SingleError(c.Writer, 0, errors.New("Falha na tentativa de enviar a requisição."))
		return nil, 0
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)

	if response.StatusCode >= http.StatusBadRequest {
		res.ClientError(c.Writer, response)
		return nil, response.StatusCode
	}
	return data, response.StatusCode
}

func parse[T any](w http.ResponseWriter, payload []byte, response T) bool {
	if err := json.Unmarshal(payload, &response); err != nil {
		res.SingleError(w, http.StatusBadRequest, errors.New("Falha ao tentar parsear a resposta da requisição."))
		return false
	}
	return true
}
