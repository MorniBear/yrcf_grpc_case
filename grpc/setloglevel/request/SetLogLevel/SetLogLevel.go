package SetLogLevel

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	setloglevel "yrcf_grpc_case/grpc/setloglevel/proto"
)

func Run(client *setloglevel.SetLogLevelClient, param *setloglevel.SetLogLevelRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).SetLogLevel(ctx, param)
	if err != nil {
		log.Fatal("did not set log level err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Set log level failed, errcode=%d, result=%s\n",
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

	client := setloglevel.NewSetLogLevelClient(conn)
	result := Run(&client, &setloglevel.SetLogLevelRequest{})
	if result {
		fmt.Println("set log level succeeded")
	}
}
