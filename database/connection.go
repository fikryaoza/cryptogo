package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var onceDb sync.Once

var instance *mongo.Client

func GetInstance() *mongo.Client {
	onceDb.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://172.17.0.1:27017")
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		err2 := client.Ping(context.Background(), readpref.Primary())
		if err2 != nil {
			log.Fatal("Couldn't connect to the database", err)
		} else {
			log.Println("Connected to Mongodb Server!")
		}
		instance = client
	})
	return instance
}
