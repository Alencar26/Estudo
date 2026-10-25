package hash

import (
	"crypto/rand"
)

const saltLength = 16

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}
