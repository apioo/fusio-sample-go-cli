package main

import (
	"fmt"
	"github.com/apioo/fusio-sdk-go"
	"github.com/apioo/fusio-sdk-go/backend"
	"github.com/apioo/sdkgen-go"
	"log"
)

func main() {
	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"backend"}

	var client = fusio.NewClient("https://demo.fusio-project.org", "test", "FRsNh1zKCXlB", store, scopes)

	collection, error := client.Backend.BackendRoute().GetBackendRoutes().BackendActionRouteGetAll(backend.CollectionCategoryQuery{})

	if error != nil {
		log.Panic(error)
	}

	for _, v := range collection.Entry {
		fmt.Println("Route: " + v.Path)
	}
}
