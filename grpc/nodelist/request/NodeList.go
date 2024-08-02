package NodeList

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	nodelist "yrcf_grpc_case/grpc/nodelist/proto"
)

func Run(client *nodelist.NodeListClient, param *nodelist.NodeListRequest) []*nodelist.NodeInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).NodeList(ctx, param)
	if err != nil {
		log.Fatal("did not get node list err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("get node list failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetNodeLists()
	}
	return nil
}

func Example(address string, port int) {
	fmt.Println("{\n  \"client\": true,\n  \"hide_internal_ips\": false,\n  \"agent\": true\n}")
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := nodelist.NewNodeListClient(conn)
	results := Run(&client, &nodelist.NodeListRequest{
		Client:          true,
		HideInternalIps: false,
		Agent:           true,
	})

	printResult(results)
}

func printResult(results []*nodelist.NodeInfo) {
	for _, result := range results {
		fmt.Printf("[%s][%s][%d]\n",
			result.GetNodeName(), result.GetType(), result.GetNodeNumId())
	}
}
