package core

import (
	"github.com/dgrijalva/jwt-go"
)

func SignJWT(headers map[string]string, claims jwt.Claims) (string, error) {
	//signingMethod := jwt.GetSigningMethod("ES256")
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//token.Header["kid"] = ""
	return "", nil
}
