package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//URL ...
type URL struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	URL       string             `bson:"url"`
	URLShort  string             `bson:"url_short"`
}

// CreateURL - Insert a new document in the collection
func CreateURL(url URL) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(DB).Collection(COLLECTION)
	_, err = collection.InsertOne(context.TODO(), url)
	if err != nil {
		return err
	}

	return nil
}

// GetURLByShort - Get URL by a short URL
func GetURLByShort(url string) (URL, error) {
	result := URL{}
	filter := bson.D{primitive.E{Key: "url_short", Value: url}}

	client, err := GetMongoClient()
	if err != nil {
		return result, err
	}

	collection := client.Database(DB).Collection(COLLECTION)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
