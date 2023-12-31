package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"twitterGo/models"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	pass := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, pass, host)

	clientOptions := options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	fmt.Println("Connection Success to DB")

	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func ConnectedDB() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
