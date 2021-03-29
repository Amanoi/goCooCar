package dao

import (
	"context"
	"os"
	"testing"

	mgo "coolcar/share/mongo"
	mongotesting "coolcar/share/mongo/testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			mgo.IDField: mustObjID("605d838cbcfcb14576815cbc"),
			opedIDField: "openid_1",
		},
		bson.M{
			mgo.IDField: mustObjID("605d838cbcfcb14576915cbe"),
			opedIDField: "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert inital values: %v", err)
	}
	m.newObjID = func() primitive.ObjectID {
		return mustObjID("605d838cbcfcb14576c15cb4")
	}
	
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "605d838cbcfcb14576815cbc",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "605d838cbcfcb14576915cbe",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "605d838cbcfcb14576c15cb4",
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("fail resolve account id for %q:%v\n", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id want: %q,got:%q", cc.want, id)
			}
		})
	}
	// id, err := m.ResolveAccountID(c, "123")
	// if err != nil {
	// 	t.Fatalf("fail resolve account id for 123: %v", err)
	// } else {
	// 	want := "605d838cbcfcb14576c15cb3"
	// 	if id != want {
	// 		t.Errorf("resolve account id want: %q,got:%q", want, id)
	// 	}
	// }
}
func mustObjID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return objID
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
