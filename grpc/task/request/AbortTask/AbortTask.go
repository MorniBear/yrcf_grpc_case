package AbortTask

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	task "yrcf_grpc_case/grpc/task/proto"
)

func Run(client *task.TaskClient, param *task.AbortTaskRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AbortTask(ctx, param)
	if err != nil {
		log.Fatal("did not abort task err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Abort task failed, errcode=%d, result=%s\n",
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

	client := task.NewTaskClient(conn)
	result := Run(&client, &task.AbortTaskRequest{
		Id:   1,
		Type: task.TaskType_kBucketlink,
	})
	if result {
		fmt.Println("Abort task succeeded")
	}
}
