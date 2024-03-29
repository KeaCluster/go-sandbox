package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID     int    `bson:"_id,omitempty" json:"id"`
	Title  string `bson:"title" json:"title"`
	Author string `bson:"author" json:"author"`
	ISBN   string `bson:"isbn" json:"isbn"`
}

// Load

func GetBooks(client *mongo.Client) ([]Book, error) {
	var books []Book
	collection := client.Database("z_lib").Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books, nil
}
