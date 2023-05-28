package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dataBase = "rarefinds"
)

var (
	Products *mongo.Collection
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//TODO: GET A MONGO URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ralkinho:0V1UfU1DYy6Za2Oi@rarefinds.e8tte3f.mongodb.net/"))
	if err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

	Products = client.Database(dataBase).Collection("products")
}