package core

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
)

func SignJWT(claims jwt.Claims, signingKey *key.SigningKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = signingKey.ID

	result, err := token.SignedString(signingKey.Key)
	if err != nil {
		return "", err
	}

	return result, nil
}
