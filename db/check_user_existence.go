package db

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UserExistenceCheck(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var user models.User

	err := col.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}
