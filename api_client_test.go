package espocrm

import (
	"os"
	"testing"
)

func TestNewApiClient(t *testing.T) {
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

// TODO: Test ApiClient.Read()

// TODO: Test ApiClient.Delete()

// TODO: Test ApiClient.List()

// TODO: Test ApiClient.Update()

// TODO: Test ApiClient.Create()
