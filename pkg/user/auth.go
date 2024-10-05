package user

import (
	"context"

	"github.com/FulgurCode/StellarPing/pkg/mongodb"
)

func (user *User) Signup() error {
	var _, err = mongodb.UserCollection().InsertOne(context.Background(), user)
	return err
}
