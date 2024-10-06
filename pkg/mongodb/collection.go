package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func UserCollection() *mongo.Collection {
	return Db.Collection("user")
}

func NewsCollection() *mongo.Collection {
	return Db.Collection("news")
}
