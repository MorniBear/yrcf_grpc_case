package DeleteTask

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	task "yrcf_grpc_case/grpc/task/proto"
)

func Run(client *task.TaskClient, param *task.DeleteTaskRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteTask(ctx, param)
	if err != nil {
		log.Fatal("did not delete task err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Delete task failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &task.DeleteTaskRequest{
		Id:   1,
		Type: task.TaskType_kBucketlink,
	})
	if result {
		fmt.Println("Delete task succeeded")
	}
}
