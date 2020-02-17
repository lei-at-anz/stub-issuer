package main

import (
	"fmt"
	"github.com/lei-at-anz/stub-issuer/pkg/core"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
	"log"
	"net/http"
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

	api, err := core.CreateAPI(signingKey)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/jwks.json", api.HandleJWKS)
	mux.HandleFunc("/jwts", api.HandleIssueJWT)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 3002),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
