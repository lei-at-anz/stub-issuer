package core

import (
	"crypto/ecdsa"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
	"github.com/lestrrat/go-jwx/jwk"
)

func CreateJWKSResponse(signingKey *key.SigningKey) (map[string]interface{}, error) {
	keyEntry, err := createJWKEntry(signingKey)
	if err != nil {
		return nil, err
	}

	jsonEntry := make(map[string]interface{})
	if err = keyEntry.PopulateMap(jsonEntry); err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["keys"] = []interface{}{jsonEntry}
	return result, nil
}

func createJWKEntry(signingKey *key.SigningKey) (jwk.Key, error) {
	privateKey := signingKey.Key.(*ecdsa.PrivateKey)
	keyEntry, err := jwk.New(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}
	if err = keyEntry.Set(jwk.KeyIDKey, signingKey.ID); err != nil {
		return nil, err
	}
	return keyEntry, nil
}
