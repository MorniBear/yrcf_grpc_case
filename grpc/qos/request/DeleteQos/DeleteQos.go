package DeleteQos

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	qos "yrcf_grpc_case/grpc/qos/proto"
)

func Run(client *qos.QosClient, param *qos.DeleteQosRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteQos(ctx, param)
	if err != nil {
		log.Fatal("did not del qos err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Del qos failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &qos.DeleteQosRequest{
		Path:      "/grpc_qos",
		Recursive: false,
		Force:     true,
	})
	if result {
		fmt.Println("Del qos succeeded")
	}
}
