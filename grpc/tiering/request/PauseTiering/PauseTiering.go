package PauseTiering

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.PauseTieringRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).PauseTiering(ctx, param)
	if err != nil {
		log.Fatal("did not pause tiering err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Pause tiering failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return false
	} else {
		return result.GetStatus()
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
	result := Run(&client, &tiering.PauseTieringRequest{
		Op: tiering.PauseTieringRequest_UPDATE,
		On: true,
	})

	fmt.Println("Pause tiering succeeded, status: ", result)
}
