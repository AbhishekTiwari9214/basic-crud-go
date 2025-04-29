package book

import (
	"basic-crud/models"
	"basic-crud/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateBook(book models.Book) (*mongo.InsertOneResult, error) {

	client, err := utils.GetMongoDbClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database(utils.GetEnv("DB_NAME")).Collection("books")

	result, err := collection.InsertOne(ctx, bson.M{
		"title":  book.Title,
		"author": book.Author,
		"year":   book.Year,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
