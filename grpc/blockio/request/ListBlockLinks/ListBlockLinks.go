package ListBlockLinks

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	blockio "yrcf_grpc_case/grpc/blockio/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *blockio.BlockIOClient, param *blockio.ListBlockLinksRequest) []*blockio.BlockLinkInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListBlockLinks(ctx, param)
	if err != nil {
		log.Fatal("did not add blockio, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("del blockio failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetBlockInfo()
	}
}

func Example(address string, port int) {
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("del not close err:", err)
		}
	}(conn)

	client := blockio.NewBlockIOClient(conn)
	var result = Run(&client, &blockio.ListBlockLinksRequest{})
	printResult(result)

}

func printResult(result []*blockio.BlockLinkInfo) {

}
