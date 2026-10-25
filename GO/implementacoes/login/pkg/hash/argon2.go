package hash

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = uint32(64 * 1024)
	parallelism = uint8(4)
	iterations  = uint32(3)
	keyLength   = uint32(32)
)

func HashPassword(password string) (hash string, err error) {

	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}

	rawHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(rawHash)

	encoded := fmt.Sprintf(
		"$argon2id$v=13$m=%d,t=%d,p=%d$%s$%s",
		memory,
		iterations,
		parallelism,
		saltB64,
		hashB64,
	)

	return encoded, nil
}

func VerifyPassword(password string, hash []byte, salt []byte) bool {
	hash2 := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)
	return bytes.Equal(hash, hash2)
}
