package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/lei-at-anz/stub-issuer/pkg/core"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
)

func main() {
	privateKey, err := key.LoadPrivateKey(&key.FileSpec{
		File:     "keys/key.pem",
		Password: "changeit",
	})
	if err != nil {
		panic(err)
	}
	writePrivateKey(privateKey.Key.(*ecdsa.PrivateKey))
}

func writePublicKey() {
	key, err := core.LoadKey()
	if err != nil {
		panic(err)
	}
	bytes, err := core.MarshalPublicKey(&key.PublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func writePrivateKey(key *ecdsa.PrivateKey) {
	bytes, err := core.MarshalPrivateKey(key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
