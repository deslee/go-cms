package data

import (
	"github.com/google/uuid"
)

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateId() string {
	newUuid, err := uuid.NewRandom()
	die(err)

	return newUuid.String()
}
