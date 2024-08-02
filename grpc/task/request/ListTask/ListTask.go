package ListTask

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	task "yrcf_grpc_case/grpc/task/proto"
)

func Run(client *task.TaskClient, param *task.ListTaskRequest) map[uint64]*task.TaskInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListTask(ctx, param)
	if err != nil {
		log.Fatal("did not list task err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List task failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetTaskMap()
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
	result := Run(&client, &task.ListTaskRequest{
		LoadId: 0,
		Type:   task.TaskType_kBucketlink,
	})

	printfResult(result)
}

func printfResult(result map[uint64]*task.TaskInfo) {}
