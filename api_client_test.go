package espocrm

import (
	"encoding/json"
	"os"
	"testing"
)

func TestApiClient(t *testing.T) {
	type Entity struct {
		Id   string
		Name string
	}
	E := &Entity{}

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

	t.Run("create new entity", func(t *testing.T) {
		url := os.Getenv("GOESPOCRM_URL")
		apiKey := os.Getenv("GOESPOCRM_APIKEY")
		entity := os.Getenv("GOESPOCRM_ENTITY")
		payload := os.Getenv("GOESPOCRM_CREATE_ENTITY_PAYLOAD")

		client := NewApiClient(url, WithApiKeyAuth(apiKey))

		result, err := client.Create(entity, payload)
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

	t.Run("update existing entity", func(t *testing.T) {
		url := os.Getenv("GOESPOCRM_URL")
		apiKey := os.Getenv("GOESPOCRM_APIKEY")
		entity := os.Getenv("GOESPOCRM_ENTITY")
		payload := os.Getenv("GOESPOCRM_UPDATE_ENTITY_PAYLOAD")

		client := NewApiClient(url, WithApiKeyAuth(apiKey))

		result, err := client.Update(entity, E.Id, payload)
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

		var updatePayload map[string]string
		err = json.Unmarshal([]byte(payload), &updatePayload)
		if err != nil {
			t.Errorf("%s", err)
		}

		if newEntity["name"] != updatePayload["name"] {
			t.Errorf("expected name {%v}, got name {%v}", updatePayload["name"], newEntity["name"])
		}
	})

	t.Run("read updated entity", func(t *testing.T) {
		url := os.Getenv("GOESPOCRM_URL")
		apiKey := os.Getenv("GOESPOCRM_APIKEY")
		entity := os.Getenv("GOESPOCRM_ENTITY")

		client := NewApiClient(url, WithApiKeyAuth(apiKey))

		result, err := client.Read(entity, E.Id)
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
		url := os.Getenv("GOESPOCRM_URL")
		apiKey := os.Getenv("GOESPOCRM_APIKEY")
		entity := os.Getenv("GOESPOCRM_ENTITY")

		client := NewApiClient(url, WithApiKeyAuth(apiKey))

		result, err := client.Delete(entity, E.Id)
		if err != nil {
			t.Errorf("%s", err)
		}

		if string(result) != "true" {
			t.Errorf("expected true, got {%v}", result)
		}
	})
}

// TODO: Test ApiClient.List()
