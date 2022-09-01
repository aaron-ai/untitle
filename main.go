package main

import (
	"context"
	"fmt"
	"log"
	"os"
	v2 "untitle/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Host       = "47.97.252.220"
	Port int32 = 8080
)

func main() {
	endpoints := fmt.Sprintf("%s:%d", Host, Port)
	conn, err := grpc.DialContext(context.TODO(), endpoints,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	conn.Connect()
	msc := v2.NewMessagingServiceClient(conn)
	req := &v2.QueryRouteRequest{}

	resp, err := msc.QueryRoute(context.TODO(), req)

	log.Printf("%v, err=%v", resp, err)
	conn.Close()
	os.Exit(1)
}
