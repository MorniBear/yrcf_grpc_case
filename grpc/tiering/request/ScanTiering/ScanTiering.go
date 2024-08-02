package ScanTiering

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.ScanTieringRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ScanTiering(ctx, param)
	if err != nil {
		log.Fatal("did not scan tiering err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Scan tiering failed, errcode=%d, result=%s\n",
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

	client := tiering.NewTieringClient(conn)
	result := Run(&client, &tiering.ScanTieringRequest{
		TieringId: 1,
	})

	if result {
		fmt.Println("Scan tiering succeeded")
	}
}
