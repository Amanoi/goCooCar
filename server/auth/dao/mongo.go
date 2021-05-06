package dao

import (
	"context"
	"fmt"

	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"

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
func NewMongo(c context.Context,db *mongo.Database) (*Mongo,error) {
	nm := mgutil.NewMongo{
		C:    c,
		DB:   db,
		Name: "account",
	}
	flag, err := nm.IsfirstCreate()
	if err != nil {
		return nil, err
	}
	mg := &Mongo{
		col: db.Collection(nm.Name),
	}
	if !flag {
		err = mg.createIndexs(c, db)
	}
	return mg, err
}

// ResolveAccountID reslove an account id from open id.
func (m *Mongo) ResolveAccountID(c context.Context, openID string) (id.AccountID, error) {
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
	return objid.ToAccountID(row.ID), nil
}

func (m *Mongo) createIndexs(c context.Context, d *mongo.Database) error {
	_, err := d.Collection("account").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{ // D 是有序键值对
			{Key: "open_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	return err
}