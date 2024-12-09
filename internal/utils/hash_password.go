package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hashedPassword), err
}

func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)) == nil
}
