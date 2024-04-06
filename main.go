package main

import (
	"fmt"
	"github.com/apioo/fusio-sdk-go/v5/sdk"
	"github.com/apioo/sdkgen-go"
	"log"
)

func main() {
	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"backend"}

	credentials := sdkgen.OAuth2{
		ClientId:         "test",
		ClientSecret:     "FRsNh1zKCXlB",
		TokenUrl:         "https://demo.fusio-project.org/authorization/token",
		AuthorizationUrl: "",
		TokenStore:       store,
		Scopes:           scopes,
	}

	var client, _ = sdk.NewClient("https://demo.fusio-project.org", credentials)

	collection, err := client.Backend().Operation().GetAll(0, 16, "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Operations:")
	for _, v := range collection.Entry {
		fmt.Println("* " + v.HttpMethod + " " + v.HttpPath)
	}
}
