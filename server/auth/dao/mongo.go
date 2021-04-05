package dao

import (
	"context"
	"fmt"

	mgutil "coolcar/shared/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const opedIDField = "open_id"

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// TODO : 同一个account最多只能又一个进行中多Trip
// TODO : 强化类型TripID
// TODO : 表格驱动测试

// NewMongo creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

// ResolveAccountID reslove an account id from open id.
func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	insertedID := mgutil.NewObjID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		opedIDField: openID,
	}, mgutil.SetOnInsert(bson.M{
		mgutil.IDFieldName: insertedID,
		opedIDField:        openID,
	}), options.
		FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	var row mgutil.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}
	return row.ID.Hex(), nil
}
