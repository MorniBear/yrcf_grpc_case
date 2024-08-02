package AddTiering

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.AddTieringRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddTiering(ctx, param)
	if err != nil {
		log.Fatal("did not add tiering err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Add tiering failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &tiering.AddTieringRequest{
		UseAbsolutePath: false,
		TieringBasicInfo: &tiering.TieringBasicInfo{
			Path:           "",
			BucketId:       0,
			BackupBucketId: 0,
			ColdTime:       0,
			FlushTimer:     "",
		},
		TieringId:     1,
		Policy:        "",
		PutThreadNum:  0,
		ScanThreadNum: 0,
	})

	if result {
		fmt.Println("add tiering succeeded")
	}
}
