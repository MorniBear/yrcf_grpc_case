package RecoverTieringStat

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.RecoverTieringStatRequest) *tiering.RecoverTieringStatInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).RecoverTieringStat(ctx, param)
	if err != nil {
		log.Fatal("did not recover tiering stat err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Recover tiering stat failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetStatInfo()
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
	result := Run(&client, &tiering.RecoverTieringStatRequest{
		TieringId: 1,
	})

	printResult(result)
}

func printResult(result *tiering.RecoverTieringStatInfo) {}
