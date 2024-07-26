package connect

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetConn(address string, port int) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", address, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect err=", err)
	}

	return conn
}
