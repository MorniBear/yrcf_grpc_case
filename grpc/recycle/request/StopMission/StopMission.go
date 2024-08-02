package StopMission

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	recycle "yrcf_grpc_case/grpc/recycle/proto"
)

func Run(client *recycle.RecycleClient, param *recycle.StopMissionRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).StopMission(ctx, param)
	if err != nil {
		log.Fatal("did not stop mission err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Stop mission failed, errcode=%d, result=%s\n",
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

	client := recycle.NewRecycleClient(conn)
	result := Run(&client, &recycle.StopMissionRequest{
		IdOfMission: 1,
	})

	if result {
		fmt.Println("Stop mission succeeded")
	}
}
