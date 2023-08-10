package jwt

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"twitterGo/models"
)

func GenerateJWT(ctx context.Context, t models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	myPassword := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"surname":   t.Surname,
		"birthday":  t.Birthday,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
