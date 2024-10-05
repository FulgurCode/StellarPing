package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt generate hash
func BcryptGenerateHash(password string) string {
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
