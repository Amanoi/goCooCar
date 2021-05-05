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
	"os"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("account1"))
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

	pm := &profileManager{}
	cm := &carMangager{}
	s := &Service{
		ProfileManager: pm,
		CarManager:     cm,
		PoiManager:     &poi.Manager{},
		Mongo:          m,
		Logger:         logger,
	}

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}

	pm.iID = "identity1"
	golden := `{"account_id":"account1","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门"},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门"},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:   "normal_create",
			tripID: "617fcb673ec1f0074e5efd81",
			want:   golden,
		},
		{
			name:       "profile_err",
			tripID:     "617fcb673ec1f0074e5efd82",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			tripID:       "617fcb673ec1f0074e5efd83",
			carVerifyErr: fmt.Errorf("carVerify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			tripID:       "617fcb673ec1f0074e5efd84",
			carUnlockErr: fmt.Errorf("carUnlock"),
			want:         golden,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifiyErr = cc.carVerifyErr
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
