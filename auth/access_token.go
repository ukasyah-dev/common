package auth

import (
	"crypto"

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

func GenerateAccessToken(privateKey crypto.PrivateKey, claims Claims) (string, error) {
	var err error

	token, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims).SignedString(privateKey)
	if err != nil {
		log.Errorf("Failed to sign jwt with claims: %s", err)
		return "", errors.Internal()
	}

	return token, nil
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
