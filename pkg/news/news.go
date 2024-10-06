package news

import (
	"context"

	"github.com/FulgurCode/StellarPing/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type News struct {
	Id          string `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Url         string `json:"url" bson:"url"`
	Image       string `json:"image" bson:"image"`
	Source      string `json:"source" bson:"source"`
	Country     string `json:"country" bson:"country"`
	Date        string `json:"published_at" bson:"published_at"`
}

func (n *News) Add() error {
	var _, err = mongodb.NewsCollection().InsertOne(context.Background(), n)
	return err
}

func (n *News) Remove() error {
	var _, err = mongodb.NewsCollection().DeleteOne(context.Background(), bson.M{"_id": n.Id})
	return err
}

func AddNews(news []interface{}) {
	mongodb.NewsCollection().InsertMany(context.Background(), news)
}