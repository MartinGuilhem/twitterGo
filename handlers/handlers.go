package handlers

import (
	"context"
	"fmt"
	"twitterGo/jwt"
	"twitterGo/models"
	"twitterGo/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi

	r.Status = 400

	isOk, statusCode, msg, _ := validateAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg

		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		case "login":
			return routers.Login(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return routers.GetProfile(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Method Invalid"

	return r
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	fmt.Println("Processing validateAuthorization ")

	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Required Token", models.Claim{}
	}

	claim, allOK, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if err != nil {
		fmt.Println("Error ProcessToken " + err.Error())
	}

	if !allOK {
		fmt.Println("Error in token " + err.Error())
		return false, 401, msg, models.Claim{}
	}

	fmt.Println("Token OK")

	return true, 200, msg, *claim
}
