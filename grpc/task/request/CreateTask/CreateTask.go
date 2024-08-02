package CreateTask

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	task "yrcf_grpc_case/grpc/task/proto"
)

func Run(client *task.TaskClient, param *task.CreateTaskRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).CreateTask(ctx, param)
	if err != nil {
		log.Fatal("did not create task err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Create task failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &task.CreateTaskRequest{
		LoadId:       1,
		Type:         task.TaskType_kBucketlink,
		Op:           task.TaskOp_kPreLoad,
		Path:         "/grpc_task",
		UseMountPath: false,
		Threads:      2,
	})
	if result {
		fmt.Println("Create task succeeded")
	}
}
