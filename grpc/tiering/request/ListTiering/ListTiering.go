package ListTiering

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.ListTieringRequest) *tiering.ListTieringResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListTiering(ctx, param)
	if err != nil {
		log.Fatal("did not stat tiering err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Stat tiering failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result
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
	result := Run(&client, &tiering.ListTieringRequest{
		TieringId:  1,
		ShowPolicy: true,
		ShowLevel:  true,
	})

	printResult(result.GetTieringInfo(), result.GetTieringLevelInfo())
}

func printResult(info []*tiering.TieringInfo, levelInfo []*tiering.TieringLevelInfo) {}
