package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func MakeHashed(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassStr := string(hashedPass)

	return hashedPassStr, nil
}

func CompareHashPass(hashedPass, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
