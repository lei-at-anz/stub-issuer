package main

import (
	"encoding/json"
	"fmt"
	"github.com/lei-at-anz/stub-issuer/pkg/core"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
	"log"
)

func main() {
	config := &key.FileSpec{
		File:     "keys/key.pem",
		Password: "changeit",
	}

	signingKey, err := key.LoadPrivateKey(config)
	if err != nil {
		log.Fatal(err)
	}

	jwkResponse, err := core.CreateJWKSResponse(signingKey)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := json.Marshal(jwkResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(raw))
}
