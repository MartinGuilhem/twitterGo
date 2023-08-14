package routers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"twitterGo/db"
	"twitterGo/models"
)

func GetProfile(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Get Profile")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID field required"
		return r
	}

	perfil, err := db.GetProfile(ID)
	if err != nil {
		r.Message = "routers.db.GetProfile " + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "routers.db.Marshal " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
