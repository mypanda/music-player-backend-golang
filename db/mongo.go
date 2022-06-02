package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	option := options.Client().ApplyURI("mongodb://music_admin:123456@localhost:27017/music_player_db")
	client, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	return client
}

// var DB *mongo.Client = ConnectDB()

func GetTestCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("test").Collection(collectionName)
	return collection
}
