package ListRecycle

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	recycle "yrcf_grpc_case/grpc/recycle/proto"
)

func Run(client *recycle.RecycleClient, param *recycle.ListRecycleRequest) []*recycle.RecycleInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListRecycle(ctx, param)
	if err != nil {
		log.Fatal("did not list recycle err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List recycle failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.RecycleInfo
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

	client := recycle.NewRecycleClient(conn)
	result := Run(&client, &recycle.ListRecycleRequest{})

	printResult(result)
}

func printResult(result []*recycle.RecycleInfo) {}
