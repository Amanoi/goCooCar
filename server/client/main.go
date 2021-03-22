package main

import (
	"context"
	"fmt"
	"log"
	trippb "server/proto/gen/go"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}
	defer conn.Close()
	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})
	if err != nil {
		log.Fatalf("Cannot call GetTrip :%v", err)
	}
	fmt.Println(r)

}
