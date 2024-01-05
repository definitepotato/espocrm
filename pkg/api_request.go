package espocrm

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func NewRequest(client *ApiClient, method, endpoint string) (*http.Request, error) {
	request, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	err = AddAuthToRequest(client, request)
	if err != nil {
		return nil, err
	}

	return request, err
}

func NewRequestWithBody(client *ApiClient, method, endpoint string, body []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	err = AddAuthToRequest(client, request)
	if err != nil {
		return nil, err
	}

	return request, err
}

func AddAuthToRequest(client *ApiClient, request *http.Request) error {
	if client.config.apiKey != nil {
		request.Header.Set("X-Api-Key", *client.config.apiKey)
		return nil
	}

	if client.config.username != nil && client.config.password != nil {
		usernameAndPassword := []byte(*client.config.username + ":" + *client.config.password)
		base64EncodedAuth := base64.StdEncoding.EncodeToString(usernameAndPassword)
		request.Header.Set("Authorization", "Basic "+base64EncodedAuth)
		return nil
	}

	return fmt.Errorf("client missing credentials")
}

func SendRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("http_status_code: %d, error: %s", response.StatusCode, response.Header.Get("X-Status-Reason"))
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, err
}
