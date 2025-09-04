package security

import (
	"github.com/csvitor-dev/go.dev/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Cryptify: hashes the password using a secure hashing algorithm
func Cryptify(password string) (string, error) {
	if password == "" {
		return "", errors.ErrPasswordNoProvided
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.ErrHashPassword
	}
	return string(hash), nil
}

// VerifyPassword: compares the hashed password with the provided password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
