package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/client/profile"
	"coolcar/rental/trip/dao"
	"coolcar/shared/server"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//DBNAME set the database name.
const DBNAME = "coolcar"

func main() {
	logger, err := server.NewZaplogger()
	if err != nil {
		log.Fatalf("cannot create logger :%v", err)
	}
	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:55000/?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect mongodb!", zap.Error(err))
	}

	db := mongoClient.Database(DBNAME)
	m, err := dao.NewMongo(c, db)
	if err != nil {
		logger.Fatal("cannot created the Mongo instance:", zap.Error(err))
	}

	logger.Sugar().Fatal(
		server.RunGRPCServer(&server.GRPCConfig{
			Name:              "rental",
			Addr:              ":8082",
			AuthPublicKeyFile: "shared/auth/public.key",
			Logger:            logger,
			RegisterFunc: func(s *grpc.Server) {
				rentalpb.RegisterTripServiceServer(s, &trip.Service{
					CarManager:     &car.Manager{},
					PoiManager:     &poi.Manager{},
					ProfileManager: &profile.Manager{},
					Mongo:          m,
					Logger:         logger,
				})
			},
		}),
	)
}
