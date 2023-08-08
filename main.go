// main.go
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
	"twitterGo/awsgo"
	"twitterGo/db"
	"twitterGo/handlers"
	"twitterGo/models"
	"twitterGo/secretmanager"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(LambdaExecute)
}

func LambdaExecute(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializeAWS()

	if !ParamValidator() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Error: env variables required: 'SecretName', 'BucketName', 'UrlPrefix'",
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Error while reading Secret " + err.Error(),
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Connect DB or Check connection
	err = db.ConnectDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Error while connecting DB " + err.Error(),
		}
		return res, nil
	}

	respAPI := handlers.Handlers(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	return respAPI.CustomResp, nil
}

func ParamValidator() bool {
	// Its mandatory that my lambda receives 3 environment variables.
	_, gotParam := os.LookupEnv("SecretName")
	if !gotParam {
		return gotParam
	}

	_, gotParam = os.LookupEnv("BucketName")
	if !gotParam {
		return gotParam
	}

	_, gotParam = os.LookupEnv("UrlPrefix")
	if !gotParam {
		return gotParam
	}

	return true
}
