package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"twitterGo/models"
)

var (
	Email  string
	IDUser string
)

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myPassword := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})
	if err == nil {
		//Routine that checks from DB
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invalid Token")
	}

	return &claims, false, string(""), err
}
