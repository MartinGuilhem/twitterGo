package routers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
	"twitterGo/db"
	"twitterGo/jwt"
	"twitterGo/models"
)

func Login(ctx context.Context) models.RespApi {
	var (
		t models.User
		r models.RespApi
	)

	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Invalid User or Password " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Email field required"
		return r
	}

	userData, exists := db.Login(t.Email, t.Password)
	if !exists {
		r.Message = "Invalid User or Password "
		return r
	}

	jwtKey, err := jwt.GenerateJWT(ctx, userData)
	if err != nil {
		r.Message = "routers.Login.GenerateJWT > " + err.Error()
		return r
	}

	resp := models.RespLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "routers.Login.json.Marshal > " + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
