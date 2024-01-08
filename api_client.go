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
	endpoint := fmt.Sprintf("%s/%s", client.url.String(), entityType)

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := NewRequestWithBody(client, method, endpoint, json.RawMessage(body))
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

	var body map[string]any
	err := json.Unmarshal([]byte(payload), &body)
	if err != nil {
		return nil, err
	}

	newBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := NewRequestWithBody(client, method, endpoint, newBody)
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
