package espocrm

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const ApiPath = "/api/v1"

type ApiClient struct {
	config Config
	url    *url.URL
}

type Config struct {
	apiKey   *string
	username *string
	password *string
}

type ApiClientOption func(c *ApiClient)

// TODO: WithHmacAuth

func WithApiKeyAuth(key string) ApiClientOption {
	return func(client *ApiClient) {
		client.config.apiKey = &key
	}
}

func WithBasicAuth(username, password string) ApiClientOption {
	return func(client *ApiClient) {
		client.config.username = &username
		client.config.password = &password
	}
}

func NewApiClient(uri string, opts ...ApiClientOption) *ApiClient {
	parsedUrl, _ := url.Parse(uri + ApiPath)

	client := &ApiClient{
		url: parsedUrl,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (client *ApiClient) Read(entityType, id string) ([]byte, error) {
	method := "GET"
	endpoint := fmt.Sprintf("%s/%s/%s", client.url.String(), entityType, id)

	request, err := NewRequest(client, method, endpoint)
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *ApiClient) Delete(entityType, id string) ([]byte, error) {
	method := "DELETE"
	endpoint := fmt.Sprintf("%s/%s/%s", client.url.String(), entityType, id)

	request, err := NewRequest(client, method, endpoint)
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *ApiClient) List(entityType string, params *Parameters) ([]byte, error) {
	method := "GET"
	client.url.Path += fmt.Sprintf("/%s", entityType)
	paramValues := url.Values{}

	if params.MaxSize != nil {
		maxSize := fmt.Sprintf("%d", params.MaxSize)
		paramValues.Add("maxSize", maxSize)
	}

	if params.OrderBy != nil {
		paramValues.Add("orderBy", *params.OrderBy)
	}

	if params.Select != nil {
		paramValues.Add("select", *params.Select)
	}

	if params.Order != nil {
		order := "desc"

		if *params.Order == Ascending {
			order = "asc"
		}
		paramValues.Add("order", order)
	}

	if params.Offset != nil {
		offset := fmt.Sprintf("%d", params.Offset)
		paramValues.Add("offset", offset)
	}

	if params.ReturnTotal != nil {
		returnTotal := "false"

		if *params.ReturnTotal {
			returnTotal = "true"
		}
		paramValues.Add("returnTotal", returnTotal)
	}

	if params.Where != nil {
		for i := 0; i < len(params.Where); i++ {
			paramValues.Add("where[%d][type]=%s", string(params.Where[i].Type))
			paramValues.Add("where[%d][attribute]=%s", params.Where[i].Attribute)
			paramValues.Add("where[%d][value]=%s", params.Where[i].Value)
		}
	}

	client.url.RawQuery = paramValues.Encode()

	request, err := NewRequest(client, method, client.url.String())
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *ApiClient) Update(entityType, id, payload string) ([]byte, error) {
	method := "PUT"
	endpoint := fmt.Sprintf("%s/%s/%s", client.url.String(), entityType, id)

	request, err := NewRequestWithBody(client, method, endpoint, json.RawMessage(payload))
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *ApiClient) Create(entityType, payload string) ([]byte, error) {
	method := "POST"
	endpoint := fmt.Sprintf("%s/%s", client.url.String(), entityType)

	request, err := NewRequestWithBody(client, method, endpoint, json.RawMessage(payload))
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}
