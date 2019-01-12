package data

import "github.com/google/uuid"

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func generateId() string {
	newUuid, err := uuid.NewRandom()
	die(err)

	return newUuid.String()
}
