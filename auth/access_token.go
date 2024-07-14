package auth

import (
	"crypto"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

type Claims struct {
	UserID     string `json:"uid,omitempty"`
	SessionID  string `json:"sid,omitempty"`
	SuperAdmin bool   `json:"adm,omitempty"`
	jwt.RegisteredClaims
}

func ParsePrivateKeyFromBase64(s string) (crypto.PrivateKey, error) {
	pem, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Errorf("Failed to decode private key: %s", err)
		return nil, errors.Internal()
	}

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM(pem)
	if err != nil {
		log.Errorf("Failed to parse private key from pem: %s", err)
		return nil, errors.Internal()
	}

	return privateKey, nil
}

func GenerateAccessToken(privateKey crypto.PrivateKey, claims Claims) (string, error) {
	var err error

	token, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims).SignedString(privateKey)
	if err != nil {
		log.Errorf("Failed to sign jwt with claims: %s", err)
		return "", errors.Internal()
	}

	return token, nil
}
