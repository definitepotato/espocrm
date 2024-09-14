## Getting Started

This Go `espocrm` package provides an API client for EspoCRM. To get started you'll have to provide the URL where EspoCRM is located and your method of authentication. Read more from the [official documentation](https://docs.espocrm.com/development/api/#authentication).

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

### Making a List request:

```go
import "github.com/definitepotato/espocrm"

parameters := espocrm.NewParameters(
    espocrm.SetWhere([]espocrm.Where{
        {
            Type: espocrm.Equals,
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

### Making a Read request:

```go
import "github.com/definitepotato/espocrm"

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

contact, err := client.Read("Contact", "78abc123def456")
```

### Making a Create request:

```go
import "github.com/definitepotato/espocrm"

newContact := `"{ "name": "Test", "assignedUserId": "1" }"`

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Create("Contact", newContact)
```

### Making an Update request:

```go
import "github.com/definitepotato/espocrm"

contactID := "993bdd81479dff4"
updatePayload := `"{ assignedUserId": "1" }"`

client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Update("Contact", contactID, updatePayload)
```

### Making a Delete request:

```go
import "github.com/definitepotato/espocrm"

contactID := "78abc123def456"
client := espocrm.NewApiClient(
    "https://espocrm.example.com",
    espocrm.WithApiKeyAuth("Your API Key here"),
)

status, err := client.Delete("Contact", contactID)
```
