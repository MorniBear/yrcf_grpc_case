package ListQos

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	qos "yrcf_grpc_case/grpc/qos/proto"
)

func Run(client *qos.QosClient, param *qos.ListQosRequest) []*qos.QosInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListQos(ctx, param)
	if err != nil {
		log.Fatal("did not list qos err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List qos failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.QosInfoList
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
	result := Run(&client, &qos.ListQosRequest{
		Path:   "/",
		Prefix: false,
	})

	printResult(result)
}

func printResult(result []*qos.QosInfo) {

}
