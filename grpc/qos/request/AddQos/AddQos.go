package AddQos

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	qos "yrcf_grpc_case/grpc/qos/proto"
)

func Run(client *qos.QosClient, param *qos.AddQosRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddQos(ctx, param)
	if err != nil {
		log.Fatal("did not add qos err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Add qos failed, errcode=%d, result=%s\n",
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

	client := qos.NewQosClient(conn)
	result := Run(&client, &qos.AddQosRequest{
		Rbps:  1000,
		Wbps:  1000,
		Riops: 1000,
		Wiops: 1000,
		Tbps:  0,
		Tiops: 0,
		Mops:  1000,
		Path:  "",
	})
	if result {
		fmt.Println("add qos succeeded")
	}
}
