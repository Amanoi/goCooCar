package mgutil

import (
	"context"
	"coolcar/shared/mongo/objid"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Common Field names.
const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

// IDField defines the object id field
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

// UpdatedAtField defines the updatedat field.
type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

// NewMongo param struct.
type NewMongo struct {
	C    context.Context
	DB   *mongo.Database
	Name string
}

//IsfirstCreate is
func (nm *NewMongo) IsfirstCreate() (bool, error) {
	flag := false

	colSlice, err := nm.DB.ListCollectionNames(nm.C, bson.M{}, nil)
	if err != nil {
		return flag, err
	}
	for _, name := range colSlice {
		if name == nm.Name {
			flag = true
			break
		}
	}
	return flag, nil
}

// NewObjID generates a new object id.
var NewObjID = primitive.NewObjectID

// NewObjectIDWithValue sets id for next objectID generation.
func NewObjectIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

// UpdatedAt  returns a value suitable for UpdateAt field.
var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a $set update doc.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert returns a $setOnInsert update doc.
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
