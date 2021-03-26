package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:55044/?readPreference=primary&ssl=false"))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Fatalf("fail resolve account id for 123: %v", err)
	} else {
		want := "605d838cbcfcb14576c15cb3"
		if id != want {
			t.Errorf("resolve account id want: %q,got:%q", want, id)
		}
	}
}
