package main

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var err error

	client := connect()
	defer client.Disconnect(context.TODO())

	sc, _ := bsoncodec.NewStructCodec(bsoncodec.DefaultStructTagParser)
	reg := bson.NewRegistryBuilder().RegisterEncoder(reflect.TypeOf(StructA{}), sc).Build()
	database := client.Database("DatabaseFoo", &options.DatabaseOptions{
		Registry: reg,
	})

	collection := database.Collection("C")
	collection.Drop(context.TODO())

	_, err = collection.InsertOne(
		context.TODO(),
		&StructA{
			ID:  ID("StructA_001"),
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
