package Rename

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	rename "yrcf_grpc_case/grpc/rename/proto"
)

func Run(client *rename.RenameClient, param *rename.RenameRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).Rename(ctx, param)
	if err != nil {
		log.Fatal("did not rename, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("rename failed, errcode=%d, result=%s\n",
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

	client := rename.NewRenameClient(conn)
	var result = Run(&client, &rename.RenameRequest{
		FromPath:     "grpc_from/grpc_file",
		ToPath:       "grpc_to/grpc_rename",
		UseMountPath: false,
	})
	if result {
		fmt.Println("rename succeeded")
	}
}
