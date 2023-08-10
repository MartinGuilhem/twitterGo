package db

import (
	"golang.org/x/crypto/bcrypt"
	"twitterGo/models"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := UserExistenceCheck(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
