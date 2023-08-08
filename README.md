
Fusio Go CLI sample
=====

# About

This is a simple Go CLI application which shows how to use the Go SDK to access a Fusio instance.
In this example we simply output all registered routes.
Fusio is an open source API management which helps to build and manage great APIs more information at:
https://www.fusio-project.org/

```go
package main

import (
    "fmt"
    "github.com/apioo/fusio-sdk-go"
    "github.com/apioo/sdkgen-go"
    "log"
)

func main() {
    var store = &sdkgen.MemoryTokenStore{}
    var scopes = []string{"backend"}

    var client = fusio.NewClient("https://demo.fusio-project.org", "test", "FRsNh1zKCXlB", store, scopes)

    backend, err := client.Backend()
    if err != nil {
        log.Panic(err)
    }

    collection, err := backend.Operation().GetAll(0, 16, "")
    if err != nil {
        log.Panic(err)
    }

    fmt.Println("Operations:")
    for _, v := range collection.Entry {
        fmt.Println("* " + v.HttpMethod + " " + v.HttpPath)
    }
}

```

# Usage

To run this app simply execute following command:

```
go run .
```
