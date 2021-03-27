package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField is
const IDField = "_id"

// ObjID defines the object id field
var ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Set returns a $set update doc.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}
