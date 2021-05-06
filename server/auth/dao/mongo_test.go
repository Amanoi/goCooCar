package dao

import (
	"context"
	"os"
	"testing"

	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"

	"go.mongodb.org/mongo-driver/bson"
)

const DbName = "coolcar"

func TestNewMongo(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	db := mc.Database(DbName)
	m, err := NewMongo(c, db)
	if err != nil {
		t.Fatalf("cannot created the Mongo instance: %v", err)
	}
	wantIndexStrings := []string{`{"_id": {"$numberInt":"1"}}`, `{"open_id": {"$numberInt":"1"}}`}
	indexSlice, err := m.col.Indexes().ListSpecifications(c)
	if indexSlice == nil {
		t.Fatalf("not found indexs in the created mongo collection")
	}
	for key, index := range indexSlice {
		if wantIndexStrings[key] != index.KeysDocument.String() {
			t.Fatalf("cannot create right index with mongo want index:%v,got:%v", wantIndexStrings[key], index.KeysDocument.String())
		}
	}
}

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	m, err := NewMongo(c, mc.Database(DbName))
	if err != nil {
		t.Fatalf("cannot created the Mongo instance: %v", err)
	}
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.AccountID("605d838cbcfcb14576815cbc")),
			opedIDField:        "openid_1",
		},
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.AccountID("605d838cbcfcb14576915cbe")),
			opedIDField:        "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert inital values: %v", err)
	}

	mgutil.NewObjectIDWithValue(id.AccountID("605d838cbcfcb14576c15cb4"))

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
			if id.String() != cc.want {
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
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
