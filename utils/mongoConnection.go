package utils

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const timeout = 10 * time.Second

func GetMongoDbClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		uri := GetEnv("MONGO_URI")
		if uri == "" {
			log.Fatal("No ENV is loaded")
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
			return
		}
		if err := client.Ping(ctx, nil); err != nil {
			clientInstanceError = err
			return
		}
		fmt.Println("✅ Connected to MongoDB!")
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}


func DisconnectDB(client *mongo.Client) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := client.Disconnect(ctx); err != nil {
        log.Println("Error while disconnecting from MongoDB:", err)
    } else {
        log.Println("Disconnected from MongoDB successfully.")
    }
}