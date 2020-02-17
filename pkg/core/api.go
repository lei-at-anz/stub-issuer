package core

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/lei-at-anz/stub-issuer/pkg/key"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	defaultSigningKey *key.SigningKey
	jwksResponse      map[string]interface{}
}

func CreateAPI(defaultSigningKey *key.SigningKey) (*API, error) {
	jwksResponse, err := CreateJWKSResponse(defaultSigningKey)
	if err != nil {
		return nil, err
	}
	return &API{
		defaultSigningKey: defaultSigningKey,
		jwksResponse:      jwksResponse,
	}, nil
}

func (api *API) HandleJWKS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	raw, err := json.Marshal(api.jwksResponse)
	if err != nil {
		internalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(raw); err != nil {
		log.Println(err)
	}
}

func (api *API) HandleIssueJWT(w http.ResponseWriter, r *http.Request) {
	log.Println("handling issue jwt ...")
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		internalServerError(w, err)
		return
	}
	claims := jwt.MapClaims{}
	if err = json.Unmarshal(payload, &claims); err != nil {
		internalServerError(w, err)
		return
	}
	if err = claims.Valid(); err != nil {
		internalServerError(w, err)
		return
	}
	token, err := SignJWT(claims, api.defaultSigningKey)
	if err != nil {
		internalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte(token)); err != nil {
		log.Println(err)
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
