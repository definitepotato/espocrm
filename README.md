## Getting Started

The `go-espocrm` package provides an API client for EspoCRM. To get started you'll have to provide the URL where EspoCRM is located and your method of authentication. Read more from the [official documentation](https://docs.espocrm.com/development/api/#authentication).

### Using API Key Authentication:

```go
client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)
```

### Using Basic Authentication (**this is highly discouraged**):

```go
client := NewApiClient(
    "https://espocrm.example.com",
    WithBasicAuth("username", "password"),
)
```

### Making a list GET request:

```go
parameters := NewParameters(
    SetWhere([]Where{
        {
            Type: Equals,
            Attribute: "myAttribute",
            Value: "myValue",
        },
    }),
)

client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)

contacts, err := client.List("Contact", parameters)
```

### Making a read GET request:

```go
client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)

contact, err := client.Read("Contact", "78abc123def456")
```

### Making a create POST request:

```go
newContact := `"{ "name": "Test", "assignedUserId": "1" }"`

client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Create("Contact", newContact)
```

### Making an update PUT request:

```go
updatePayload := `"{ assignedUserId": "1" }"`

client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)

attributes, err := client.Update("Contact", updatePayload)
```

### Making a delete DELETE request:

```go
client := NewApiClient(
    "https://espocrm.example.com",
    WithApiKeyAuth("Your API Key here"),
)

status, err := client.Delete("Contact", "78abc123def456")
```
