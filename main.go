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

	collection, err := backend.Operation().GetAll()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Operations:")
	for _, v := range collection.Entry {
		fmt.Println("* " + v.HttpMethod + " " + v.HttpPath)
	}
}
