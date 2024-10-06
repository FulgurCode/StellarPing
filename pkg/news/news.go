package news

import (
	"context"
	"fmt"

	"github.com/FulgurCode/StellarPing/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	Id        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	News      string `json:"news_text" bson:"news_text"`
	Url       string `json:"site_url" bson:"site_ur"`
	Image     string `json:"image_url" bson:"image_url"`
	NewsShort string `json:"news_summary_short" bson:"news_summary_short"`
	NewsLong  string `json:"news_summary_long" bson:"news_summary_long"`
	Time      string `json:"timestamp" bson:"timestamp"`
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

func GetNews() []News {
	var cursor, err = mongodb.NewsCollection().Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var allNews []News

	for cursor.Next(context.Background()) {
		var n News
		cursor.Decode(&n)
		allNews = append(allNews, n)
	}

	return allNews
}

func GetOneNews(id string) News {
	var news News
	var objectId, _ = primitive.ObjectIDFromHex(id)

	var result = mongodb.NewsCollection().FindOne(context.Background(), bson.M{"_id": objectId})
	result.Decode(&news)

	return news
}
