package id

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ukasyah-dev/common/log"
)

func New(l ...int) string {
	length := 22
	if len(l) > 0 {
		length = l[0]
	}

	id, err := gonanoid.New(length)
	if err != nil {
		log.Panicf("Failed to generate id: %s", err)
	}

	return id
}
