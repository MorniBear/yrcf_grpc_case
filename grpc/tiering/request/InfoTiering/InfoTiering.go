package InfoTiering

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	tiering "yrcf_grpc_case/grpc/tiering/proto"
)

func Run(client *tiering.TieringClient, param *tiering.InfoTieringRequest) []*tiering.TieringRecord {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).InfoTiering(ctx, param)
	if err != nil {
		log.Fatal("did not recover tiering stat err: ", err)
	}

	if result.GetResult() != 0 {
		log.Printf("Recover tiering stat failed, errcode=%d",
			result.GetResult())
		return nil
	} else {
		return result.GetRecords()
	}
	return nil
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
	result := Run(&client, &tiering.InfoTieringRequest{
		TieringId: 1,
		Op:        tiering.InfoTieringRequest_GET_CURRENT,
	})

	printResult(result)
}

func printResult(result []*tiering.TieringRecord) {}
