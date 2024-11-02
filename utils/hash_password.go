package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	setPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	return string(setPassword)
}
