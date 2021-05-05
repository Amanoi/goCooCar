package dao

import (
	"context"
	"fmt"
	"os"
	"testing"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
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
	wantIndexStrings := []string{`{"_id": {"$numberInt":"1"}}`, `{"trip.accountid": {"$numberInt":"1"},"trip.status": {"$numberInt":"1"}}`}
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

func TestCreateTrip(t *testing.T) {
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

	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			tripID:     "607fcb673ec1f0074e5efc96",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FNISHED,
		},
		{
			name:       "finished",
			tripID:     "607fcb673ec1f0074e5efc97",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FNISHED,
		},
		{
			name:       "in_progress",
			tripID:     "607fcb673ec1f0074e5efc99",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "in_progress",
			tripID:     "607fcb673ec1f0074e5efc00",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "anther_in_progress",
			tripID:     "607fcb673ec1f0074e5efd91",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	for _, cc := range cases {
		mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s:error expected;got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s:error creating trip: %v", cc.name, err)
			continue
		}
		if tr.ID.Hex() != cc.tripID {
			t.Errorf("%s:incorrect trip id; want:%q;got %q", cc.name, cc.tripID, tr.ID.Hex())
		}
	}

}

func TestGetTrip(t *testing.T) {
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

	acct := id.AccountID("account2")
	mgutil.NewObjID = primitive.NewObjectID
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "start",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "end",
			FeeCent:  10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
		},
		Status: rentalpb.TripStatus_FNISHED,
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	got, err := m.GetTrip(c, objid.ToTripID(tr.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}
	if diff := cmp.Diff(tr, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs; -want +got: %s", diff)
	}
}

func TestGetTrips(t *testing.T) {
	rows := []struct {
		id        string
		accountID string
		status    rentalpb.TripStatus
	}{
		{
			id:        "607fcb673ec1f0074e5efd81",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FNISHED,
		},
		{
			id:        "607fcb673ec1f0074e5efd82",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FNISHED,
		},
		{
			id:        "607fcb673ec1f0074e5efd83",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FNISHED,
		},
		{
			id:        "607fcb673ec1f0074e5efd84",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "607fcb673ec1f0074e5efd85",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	db := mc.Database(DbName)
	m, err := NewMongo(c, db)
	if err != nil {
		t.Fatalf("cannot created the Mongo instance: %v", err)
	}

	for _, r := range rows {
		mgutil.NewObjectIDWithValue(id.TripID(r.id))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: r.accountID,
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCont   int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCont:  4,
		},
		{
			name:       "get_in_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCont:   1,
			wantOnlyID: "607fcb673ec1f0074e5efd84",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(), id.AccountID(cc.accountID), cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}
			if cc.wantCont != len(res) {
				t.Errorf("incorrect result count; want: %d,got: %d", cc.wantCont, len(res))
			}
			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("only_id incorrect; want: %q, got: %q", cc.wantOnlyID, res[0].ID.Hex())
				}
			}

		})
	}
}

func TestUpateTrip(t *testing.T) {
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

	tid := id.TripID("617fcb673ec1f0074e5efd81")
	aid := id.AccountID("account_for_update")

	var now int64 = 10000
	mgutil.NewObjectIDWithValue(tid)
	mgutil.UpdatedAt = func() int64 {
		return now
	}

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	if tr.UpdatedAt != 10000 {
		t.Fatalf("wrong updateat: want: 10000,got: %d", tr.UpdatedAt)
	}
	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	}
	cases := []struct {
		name          string
		now           int64
		withUpdatedAt int64
		wantErr       bool
	}{
		{
			name:          "normal_update",
			now:           20000,
			withUpdatedAt: 10000,
		},
		{
			name:          "update_with_stale_timestamp",
			now:           30000,
			withUpdatedAt: 10000,
			wantErr:       true,
		},
		{
			name:          "update_with_refetch",
			now:           40000,
			withUpdatedAt: 20000,
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdatedAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		} else {
			if err != nil {
				t.Errorf("%s: cannot update: %v", cc.name, err)
			}
		}
		updatedTrip, err := m.GetTrip(c, tid, aid)
		fmt.Printf("updatedTrip: %+v", updatedTrip)
		if err != nil {
			t.Errorf("%s: cannot get tirp after update: %v", cc.name, err)
		}
		if cc.now != updatedTrip.UpdatedAt {
			t.Errorf("%s: incorrect updatedat: want %d,got %d", cc.name, cc.now, updatedTrip.UpdatedAt)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
