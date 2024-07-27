package auth

import (
	"crypto"
	"encoding/base64"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/golang-jwt/jwt/v5"
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

func ParsePublicKeyFromBase64(s string) (crypto.PublicKey, error) {
	pem, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Errorf("Failed to decode public key: %s", err)
		return nil, errors.Internal()
	}

	publicKey, err := jwt.ParseEdPublicKeyFromPEM(pem)
	if err != nil {
		log.Errorf("Failed to parse public key from pem: %s", err)
		return nil, errors.Internal()
	}

	return publicKey, nil
}

func VerifyAccessToken(publicKey crypto.PublicKey, accessToken string) (*Claims, error) {
	parsed, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		log.Debugf("Failed to parse with claims: %s", err)
		return nil, errors.Internal()
	}

	if claims, ok := parsed.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, errors.Internal("Unknown claims type")
}
