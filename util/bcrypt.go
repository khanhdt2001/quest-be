package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(input string) (string, error) {

	output, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing string: %v", err)
	}
	return string(output), nil
}

func CompareHashAndPassword(hashed string, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}
