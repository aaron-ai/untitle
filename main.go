package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"strings"
	v2 "untitle/v2"

	"google.golang.org/grpc"
)

var (
	//Host       = "47.97.252.220"
	//Port int32 = 8081

	Host       = "127.0.0.1"
	Port int32 = 9876
)

func main() {
	endpoints := fmt.Sprintf("%s:%d", Host, Port)
	for true {
		var dopts []grpc.DialOption
		newTLS := credentials.NewTLS(&tls.Config{
			RootCAs:            x509.NewCertPool(),
			InsecureSkipVerify: true,
		})
		dopts = append(dopts, grpc.WithTransportCredentials(newTLS))
		//dopts = append(dopts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.DialContext(context.TODO(), endpoints, dopts...)
		if err != nil {
			panic(err)
		}
		conn.Connect()
		msc := v2.NewMessagingServiceClient(conn)
		req := &v2.QueryRouteRequest{}

		resp, err := msc.QueryRoute(context.TODO(), req)

		if nil != err {
			log.Printf("%v, err=%v", resp, err)
			if !strings.Contains(err.Error(), "empty") && !strings.Contains(err.Error(), "Unknown") {
				break
			}
		} else {
			log.Print(resp)
		}

		conn.Close()
		//os.Exit(1)
	}
}
