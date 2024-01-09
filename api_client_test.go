package espocrm

// Set the  env variables below prior to running tests, we're making
// API calls here to make sure things work, no mocking. Run against a test
// EspoCRM instance. The values below are examples
//
// GOESPOCRM_URL="https://address-of-your-espocrm"
// GOESPOCRM_APIKEY="abc123def456"
// GOESPOCRM_USERNAME="user"
// GOESPOCRM_PASSWORD="pass"
// GOESPOCRM_ENTITY="Contact"
// GOESPOCRM_CREATE_ENTITY_PAYLOAD='{ "name": "Alice" }'
// GOESPOCRM_UPDATE_ENTITY_PAYLOAD='{ "name": "Bob" }'

import (
	"encoding/json"
	"os"
	"testing"
)

func TestApiClient(t *testing.T) {
	t.Run("new api client with api key", func(t *testing.T) {
		url := os.Getenv("GOESPOCRM_URL")
		apiKey := os.Getenv("GOESPOCRM_APIKEY")

		client := NewApiClient(url, WithApiKeyAuth(apiKey))

		if *client.config.apiKey != apiKey {
			t.Errorf("got {%v}, wanted {%v}", *client.config.apiKey, apiKey)
		}
	})

	t.Run("new api client with basic auth", func(t *testing.T) {
		url := os.Getenv("GOESPOCRM_URL")
		username := os.Getenv("GOESPOCRM_USERNAME")
		password := os.Getenv("GOESPOCRM_PASSWORD")

		client := NewApiClient(url, WithBasicAuth(username, password))

		if *client.config.username != username {
			t.Errorf("got username {%v}, wanted {%v}", *client.config.username, username)
		}

		if *client.config.password != password {
			t.Errorf("got password {%v}, wanted {%v}", *client.config.password, password)
		}
	})
}

func TestApiClientCRUD(t *testing.T) {
	type Entity struct {
		Id   string
		Name string
	}
	E := &Entity{}

	type Client struct {
		Url     string
		ApiKey  string
		Entity  string
		Payload string
	}
	C := &Client{
		Url:    os.Getenv("GOESPOCRM_URL"),
		ApiKey: os.Getenv("GOESPOCRM_APIKEY"),
		Entity: os.Getenv("GOESPOCRM_ENTITY"),
	}

	t.Run("create new entity", func(t *testing.T) {
		payload := os.Getenv("GOESPOCRM_CREATE_ENTITY_PAYLOAD")
		endpoint := NewApiClient(C.Url, WithApiKeyAuth(C.ApiKey))

		result, err := endpoint.Create(C.Entity, payload)
		if err != nil {
			t.Errorf("%s", err)
		}

		var newEntity map[string]any
		err = json.Unmarshal(result, &newEntity)
		if err != nil {
			t.Errorf("%s", err)
		}

		if newEntity == nil {
			t.Errorf("entity is empty, expected new created entity")
		}

		E.Id = newEntity["id"].(string)
		E.Name = newEntity["name"].(string)
	})

	t.Run("update created entity", func(t *testing.T) {
		payload := os.Getenv("GOESPOCRM_UPDATE_ENTITY_PAYLOAD")
		endpoint := NewApiClient(C.Url, WithApiKeyAuth(C.ApiKey))

		result, err := endpoint.Update(C.Entity, E.Id, payload)
		if err != nil {
			t.Errorf("%s", err)
		}

		var updatedEntity map[string]any
		err = json.Unmarshal(result, &updatedEntity)
		if err != nil {
			t.Errorf("%s", err)
		}

		if updatedEntity == nil {
			t.Errorf("expected an entity")
		}

		var updatePayload map[string]any
		err = json.Unmarshal([]byte(payload), &updatePayload)
		if err != nil {
			t.Errorf("%s", err)
		}

		if updatedEntity["name"] != updatePayload["name"] {
			t.Errorf("expected name {%v}, got name {%v}", updatePayload["name"], updatedEntity["name"])
		}
	})

	t.Run("read updated entity", func(t *testing.T) {
		endpoint := NewApiClient(C.Url, WithApiKeyAuth(C.ApiKey))

		result, err := endpoint.Read(C.Entity, E.Id)
		if err != nil {
			t.Errorf("%s", err)
		}

		var updatedEntity map[string]any
		err = json.Unmarshal(result, &updatedEntity)
		if err != nil {
			t.Errorf("%s", err)
		}

		if E.Name == updatedEntity["name"].(string) {
			t.Errorf("expected entity name to change")
		}
	})

	t.Run("delete updated entity", func(t *testing.T) {
		endpoint := NewApiClient(C.Url, WithApiKeyAuth(C.ApiKey))

		result, err := endpoint.Delete(C.Entity, E.Id)
		if err != nil {
			t.Errorf("%s", err)
		}

		if string(result) != "true" {
			t.Errorf("expected true, got {%v}", result)
		}
	})
}

// TODO: Test ApiClient.List()
