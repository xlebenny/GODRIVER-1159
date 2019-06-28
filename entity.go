package main

// StructA .
type StructA struct {
	ID  `bson:"_id"`
	Foo string `bson:"Foo"`
}

// ID .
type ID string

// // MarshalBSONValue .
// func (e ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
// 	return bsontype.String, bsoncore.AppendString(nil, string(e)), nil
// }
