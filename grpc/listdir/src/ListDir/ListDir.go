package ListDir

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	listdir "yrcf_grpc_case/grpc/listdir/proto"
)

func Run(client *listdir.ListDirClient, param *listdir.ListDirRequest) []*listdir.ListDirInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListDir(ctx, param)
	if err != nil {
		log.Fatal("List Dir err:", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List Dir, errcode=%d, result=%s\n", result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	}
	return result.GetListdirInfoList()
}

func Example(addr string, port int) {
	conn := connect.GetConn(addr, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := listdir.NewListDirClient(conn)
	result := Run(&client, &listdir.ListDirRequest{
		Path:    "/",
		Mounted: false,
	})

	printResult(result)
}

func printResult(result []*listdir.ListDirInfo) {
	for _, info := range result {
		fmt.Printf("[%s][%s]\n", info.GetEntryTypes(), info.GetEntryNames())
	}
}
