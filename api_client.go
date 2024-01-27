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

// Read will fetch an `entityType` record based on an `id`.
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

// Delete will delete an `entityType` record based on an `id`.
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

	client.url.RawQuery = params.Encode()

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

// Update will update an `entityType` record based on an `id`. The `payload` should be a json string
// based on documentation found in https://docs.espocrm.com/development/api/crud/#update
// and should contain attributes to update for the given record.
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

// Create will create a new `entityType` record. The `payload` should be a json string
// based on documentation found in https://docs.espocrm.com/development/api/crud/#create
// and should contain attributes for the new record.
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

// ListRelated will fetch all records related to `entityType` based on `id`.
// The `relatedEntityType` is the related record type, based on documentation
// found in https://docs.espocrm.com/development/api/relationships/#list-related-records.
func (client *ApiClient) ListRelated(entityType, id, relatedEntityType string, params *Parameters) ([]byte, error) {
	method := "GET"
	client.url.Path += fmt.Sprintf("/%s/%s/%s", entityType, id, relatedEntityType)

	client.url.RawQuery = params.Encode()

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

// Link relates two records as documented at https://docs.espocrm.com/development/api/relationships/#link.
func (client *ApiClient) Link(entityType, entityId, relatedEntityType string, relatedEntityIds []string) ([]byte, error) {
	method := "POST"
	endpoint := fmt.Sprintf("/%s/%s/%s", entityType, entityId, relatedEntityType)

	payload := struct {
		Ids []string `json:"ids"`
	}{
		Ids: relatedEntityIds,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := NewRequestWithBody(client, method, endpoint, jsonPayload)
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Unlink unrelates two records as documented at https://docs.espocrm.com/development/api/relationships/#unlink.
func (client *ApiClient) Unlink(entityType, entityId, relatedEntityType string, relatedEntityIds []string) ([]byte, error) {
	method := "DELETE"
	endpoint := fmt.Sprintf("/%s/%s/%s", entityType, entityId, relatedEntityType)

	payload := struct {
		Ids []string `json:"ids"`
	}{
		Ids: relatedEntityIds,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := NewRequestWithBody(client, method, endpoint, jsonPayload)
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, err
}
