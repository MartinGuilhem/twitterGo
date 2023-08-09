package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"twitterGo/db"
	"twitterGo/models"
)

func Register(ctx context.Context) models.RespApi {
	var (
		t models.User
		r models.RespApi
	)

	r.Status = 400

	fmt.Println("Enter Registration")

	body := ctx.Value(models.Key("body")).(string)
	if err := json.Unmarshal([]byte(body), &t); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Email Required"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Password must have more than 6 characters"
		fmt.Println(r.Message)
		return r
	}

	_, found, _ := db.UserExistenceCheck(t.Email)
	if found {
		r.Message = "User already exists"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		r.Message = "Error while registering user " + err.Error()
		fmt.Println(r.Message)
	}

	// TODO: refactor error propagation.
	if !status {
		r.Message = "User registration failed"
	}

	r.Status = 200
	r.Message = "Registration OK"
	fmt.Println(r.Message) // TODO: refactor logging.

	return r
}
