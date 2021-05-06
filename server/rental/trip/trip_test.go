package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"coolcar/shared/server"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	pm := &profileManager{}
	cm := &carMangager{}

	s := newService(c, t, pm, cm)

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}

	pm.iID = "identity1"
	golden := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":1620295450},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":1620295450},"status":1,"identity_id":"identity1"}`
	nowFunc = func() int64 {
		return 1620295450
	}
	cases := []struct {
		name         string
		accountID    string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:      "normal_create",
			accountID: "account1",
			tripID:    "617fcb673ec1f0074e5efd81",
			want:      fmt.Sprintf(golden, "account1"),
		},
		{
			name:       "profile_err",
			accountID:  "account2",
			tripID:     "617fcb673ec1f0074e5efd82",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			accountID:    "account3",
			tripID:       "617fcb673ec1f0074e5efd83",
			carVerifyErr: fmt.Errorf("carVerify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			accountID:    "account4",
			tripID:       "617fcb673ec1f0074e5efd84",
			carUnlockErr: fmt.Errorf("carUnlock"),
			want:         fmt.Sprintf(golden, "account4"),
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifiyErr = cc.carVerifyErr
			c := auth.ContextWithAccountID(context.Background(), id.AccountID(cc.accountID))
			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error; got none")
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("error creating trip: %v", err)
				return
			}
			if res.Id != cc.tripID {
				t.Errorf("incorrect id; want: %q,got: %q", cc.tripID, res.Id)
			}
			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("cannot marshal response: %v", err)
			}
			got := string(b)
			if cc.want != got {
				t.Errorf("incorrect response: want %s,got %s", cc.want, got)
			}
		})
	}
}

func TestTripLifecycle(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("account_for_lifecycle"))
	pm := &profileManager{}
	cm := &carMangager{}
	s := newService(c, t, pm, cm)
	tid := id.TripID("627fcb673ec1f0074e5efd84")
	mgutil.NewObjectIDWithValue(tid)
	cases := []struct {
		name string
		now  int64
		op   func() (*rentalpb.Trip, error)
		want string
	}{
		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					CarId: "car1",
					Start: &rentalpb.Location{
						Latitude:  32.123,
						Longitude: 114.2525,
					}})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  33.1024,
						Longitude: 115.2213,
					},
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"location":{"latitude":33.1024,"longitude":115.2213},"fee_cent":8522,"km_driven":120.94217315616464,"poi_name":"天安门","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"location":{"latitude":33.1024,"longitude":115.2213},"fee_cent":9218,"km_driven":171.21893956988873,"poi_name":"天安门","timestamp_sec":30000},"end":{"location":{"latitude":33.1024,"longitude":115.2213},"fee_cent":9218,"km_driven":171.21893956988873,"poi_name":"天安门","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"location":{"latitude":33.1024,"longitude":115.2213},"fee_cent":9218,"km_driven":171.21893956988873,"poi_name":"天安门","timestamp_sec":30000},"end":{"location":{"latitude":33.1024,"longitude":115.2213},"fee_cent":9218,"km_driven":171.21893956988873,"poi_name":"天安门","timestamp_sec":30000},"status":2}`,
		},
	}
	rand.Seed(1234)
	for _, cc := range cases {
		nowFunc = func() int64 {
			return cc.now
		}
		trip, err := cc.op()

		if err != nil {
			t.Errorf("%s: operat failed : %v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s: failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s: incorrect response; want: %s,got: %s", cc.name, cc.want, got)
		}

	}
}

func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := server.NewZaplogger()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	db := mc.Database("coolcar")
	m, err := dao.NewMongo(c, db)
	if err != nil {
		t.Fatalf("cannot created the Mongo instance: %v", err)
	}

	return &Service{
		ProfileManager: pm,
		CarManager:     cm,
		PoiManager:     &poi.Manager{},
		Mongo:          m,
		Logger:         logger,
	}
}

type profileManager struct {
	iID id.IdentityID
	err error
}

func (p *profileManager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return p.iID, p.err
}

type carMangager struct {
	verifiyErr error
	unlockErr  error
}

func (c *carMangager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return c.verifiyErr
}

func (c *carMangager) Unlock(context.Context, id.CarID) error {
	return c.unlockErr
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
