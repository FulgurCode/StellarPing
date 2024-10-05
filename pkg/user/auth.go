package user

import (
	"context"

	"github.com/FulgurCode/StellarPing/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func (user *User) Signup() error {
	var _, err = mongodb.UserCollection().InsertOne(context.Background(), user)
	return err
}

func GetUser(email string) *User {
	var user *User
	mongodb.UserCollection().FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	return user
}
