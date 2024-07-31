package GetEntry

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	getentry "yrcf_grpc_case/grpc/getentry/proto"
)

func Run(client *getentry.GetEntryClient, param *getentry.GetEntryRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).GetEntry(ctx, param)
	if err != nil {
		log.Fatal("did not get cidr priority, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Get cidr priority failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return false
	} else {
		return true
	}
}

func Example(address string, port int) {
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := getentry.NewGetEntryClient(conn)
	result := Run(&client, &getentry.GetEntryRequest{
		UseAbsolutePath: false,
		Path:            "/testfile",
	})
	if result {
		log.Println("success")
	}
}
