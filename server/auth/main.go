package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/auth/wechat"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/shared/server"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"net/http"
// 	trippb "server/proto/gen/go"
// 	trip "server/tripservice"

// 	"github.com/grpc-ecosystem/grpc-gateway/runtime"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	log.SetFlags(log.Lshortfile)
// 	go startGRPCGateway()
// 	lis, err := net.Listen("tcp", ":8081")
// 	if err != nil {
// 		log.Fatalf("failed to listen %v", err)
// 	}
// 	defer lis.Close()
// 	s := grpc.NewServer()
// 	trippb.RegisterTripServiceServer(s, &trip.Service{})
// 	log.Fatal(s.Serve(lis))

// }
// func startGRPCGateway() {
// 	c := context.Background()
// 	c, cancel := context.WithCancel(c)
// 	defer cancel()
// 	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		EnumsAsInts: true,
// 		OrigName:    true,
// 	}))
// 	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
// 	if err != nil {
// 		log.Fatalf("Cannot start grpc gateway :%v", err)
// 	}
// 	err = http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		log.Fatalf("Cannot listen and server grpc gateway :%v", err)
// 	}
// }

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

	pkFile, err := os.Open("auth/private.key")

	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)

	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	db := mongoClient.Database(DBNAME)
	m, err := dao.NewMongo(c, db)
	if err != nil {
		logger.Fatal("cannot created the Mongo instance:", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(
		&server.GRPCConfig{
			Name:   "auth",
			Addr:   ":8081",
			Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				authpb.RegisterAuthServiceServer(s, &auth.Service{
					OpenIDResolver: &wechat.Service{
						AppID:     "wx0f0f861580b7e1d1",
						AppSecret: "0cc481da864bfeab26182d0bf82b12c5",
					},
					Mongo:          m,
					Logger:         logger,
					TokenExpire:    5 * time.Second, //.Hour,
					TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privKey),
				})
			},
		}),
	)
}
