package hash

import (
	"github.com/matthewhartstonge/argon2"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

var argon = argon2.DefaultConfig()

func Generate(s string) string {
	encoded, err := argon.HashEncoded([]byte(s))
	if err != nil {
		log.Panicf("Failed to generate hash: %s", err)
	}

	return string(encoded)
}

func Verify(s, hash string) error {
	ok, err := argon2.VerifyEncoded([]byte(s), []byte(hash))
	if err != nil {
		log.Errorf("Failed to verify hash: %s", err)
		return errors.Internal()
	}

	if !ok {
		return errors.InvalidArgument()
	}

	return nil
}