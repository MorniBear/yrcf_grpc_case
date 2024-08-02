package RecycleFile

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	recycle "yrcf_grpc_case/grpc/recycle/proto"
)

func Run(client *recycle.RecycleClient, param *recycle.RecycleFileRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).RecycleFile(ctx, param)
	if err != nil {
		log.Fatal("did not recycle file err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Recycle file failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &recycle.RecycleFileRequest{
		Id:         1,
		Path:       "/grpc_recycle/recycle_file",
		Eid:        "",
		OwnerMdsId: 1,
	})

	if result {
		fmt.Println("Recycle file succeeded")
	}
}
