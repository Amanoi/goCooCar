package dao

import (
	"context"
	"fmt"

	mgo "coolcar/share/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const opedIDField = "open_id"

// Mongo defines a mongo dao.
type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

// NewMongo creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

// ResolveAccountID reslove an account id from open id.
func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	insertedID := m.newObjID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		opedIDField: openID,
	}, mgo.SetOnInsert(bson.M{
		mgo.IDField: insertedID,
		opedIDField: openID,
	}), options.
		FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	var row = mgo.ObjID
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}
	return row.ID.Hex(), nil
}
