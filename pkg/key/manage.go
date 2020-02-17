package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
)

func LoadPrivateKey(config *FileSpec) (*SigningKey, error) {
	raw, err := ioutil.ReadFile(config.File)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(raw)
	b, err := x509.DecryptPEMBlock(block, []byte(config.Password))
	if err != nil {
		return nil, err
	}
	key, err := x509.ParseECPrivateKey(b)
	if err != nil {
		return nil, err
	}
	return &SigningKey{
		Key: key,
		ID:  block.Headers["kid"],
	}, nil
}

func GeneratePrivateKey(keyType KeyType) (*SigningKey, error) {
	result := SigningKey{
		ID: uuid.New().String(),
	}
	switch keyType {
	case KeyTypeES256:
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, err
		}
		result.Key = key
		break
	case KeyTypeES384:
		key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, err
		}
		result.Key = key
		break
	case KeyTypeES512:
		key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			return nil, err
		}
		result.Key = key
		break
	default:
		return nil, errors.New("unsupported key alg")
	}
	return &result, nil
}
