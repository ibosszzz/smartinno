package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_NAME = "go-test"

func GetDBCollection(db_name string, collection_name string) (*mongo.Collection, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database(db_name).Collection(collection_name)
	return collection, nil
}