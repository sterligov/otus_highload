package crypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(s string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt generate from password: %w", err)
	}

	return string(encrypted), nil
}
