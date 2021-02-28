package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

const (
	// CONNECTIONSTRING ...
	CONNECTIONSTRING = "mongodb://root:example@localhost:27017"
	// DB ...
	DB = "url_shortener"
	// COLLECTION ...
	COLLECTION = "urls"
)

// GetMongoClient ...
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Println("Connection error")
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Println("Ping error")
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
