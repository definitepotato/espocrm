## Getting Started

The `go-espocrm` package provides an API client for EspoCRM. To get started you'll have to provide the URL where EspoCRM is located and your method of authentication. Read more from the [official documentation](https://docs.espocrm.com/development/api/#authentication).

### Using API Key Authentication:

```go
import "github.com/definitepotato/espocrm"

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)
```

### Using Basic Authentication (**this is highly discouraged**):

```go
import "github.com/definitepotato/espocrm"

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithBasicAuth("username", "password"),
)
```

### Making a list GET request:

```go
import "github.com/definitepotato/espocrm"

parameters := espocrm.NewParameters(
    espocrm.SetWhere([]Where{
        {
            Type: Equals,
            Attribute: "myAttribute",
            Value: "myValue",
        },
    }),
)

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

contacts, err := client.List("Contact", parameters)
```

### Making a read GET request:

```go
import "github.com/definitepotato/espocrm"

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

contact, err := client.Read("Contact", "78abc123def456")
```

### Making a create POST request:

```go
import "github.com/definitepotato/espocrm"

newContact := `"{ "name": "Test", "assignedUserId": "1" }"`

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Create("Contact", newContact)
```

### Making an update PUT request:

```go
import "github.com/definitepotato/espocrm"

updatePayload := `"{ assignedUserId": "1" }"`

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Update("Contact", updatePayload)
```

### Making a delete DELETE request:

```go
import "github.com/definitepotato/espocrm"

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

status, err := client.Delete("Contact", "78abc123def456")
```
