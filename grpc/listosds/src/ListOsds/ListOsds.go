package ListOsds

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	listosds "yrcf_grpc_case/grpc/listosds/proto"
)

func Run(client *listosds.ListOsdsClient, param *listosds.ListOsdsRequest) []*listosds.OsdInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListOsds(ctx, param)
	if err != nil {
		log.Fatal("List Osds err:", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List Osds, errcode=%d, result=%s\n", result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	}
	return result.GetOsdInfoList()
}

func Example(addr string, port int) {
	conn := connect.GetConn(addr, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := listosds.NewListOsdsClient(conn)
	result := Run(&client, &listosds.ListOsdsRequest{
		OsdType: 1,
		PoolId:  "",
	})

	printResult(result)
}

func printResult(result []*listosds.OsdInfo) {
	for _, osd := range result {
		fmt.Println(osd)
	}
}
