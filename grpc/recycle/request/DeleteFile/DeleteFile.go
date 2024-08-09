package DeleteFile

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	recycle "yrcf_grpc_case/grpc/recycle/proto"
)

func Run(client *recycle.RecycleClient, param *recycle.DeleteFileRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteFile(ctx, param)
	if err != nil {
		log.Fatal("did not del file err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Del file failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &recycle.DeleteFileRequest{
		Id:         1,
		Path:       "/grpc_recycle/del_file",
		Eid:        "0012D66B4AA52",
		OwnerMdsId: 1,
	})

	if result {
		fmt.Println("Del file succeeded")
	}
}
