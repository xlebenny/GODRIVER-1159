package main

import (
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// StructA .
type StructA struct {
	ID `bson:"_id"`
	Foo string `bson:"Foo"`
}

// ID .
type ID string 

// MarshalBSONValue .
func (e ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.String, bsoncore.AppendString(nil, string(e)), nil
}
