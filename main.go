package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var err error

	client := connect()
	defer client.Disconnect(context.TODO())

	database := client.Database("DatabaseFoo")
	collection := database.Collection("CollectionFoo")
	collection.Drop(context.TODO())

	_, err = collection.InsertOne(
		context.TODO(),
		&StructA{
			ID: ID("StructA_001"),
			Foo: "Bar",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Complete")
}

func connect() *mongo.Client {
	var err error
	var client *mongo.Client

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
