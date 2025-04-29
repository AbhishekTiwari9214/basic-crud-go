package book

import (
	"basic-crud/models"
	"basic-crud/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBookInfo(id string) (*models.Book, error) {
	client, err := utils.GetMongoDbClient()
	if err != nil {
		return nil, err
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid book ID")
	}
	var book models.Book
	collection := client.Database(utils.GetEnv("DB_NAME")).Collection("books")
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&book); err != nil {
		return nil, err
	}

	
	return &book, nil

}
