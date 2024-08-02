package Mkfs

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	mkfs "yrcf_grpc_case/grpc/mkfs/proto"
)

func Run(client *mkfs.MkfsClient, param *mkfs.MkfsRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).Mkfs(ctx, param)
	if err != nil {
		log.Fatal("did not mkfs err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Mkfs failed, errcode=%d, result=%s\n",
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

	client := mkfs.NewMkfsClient(conn)
	result := Run(&client, &mkfs.MkfsRequest{})
	if result {
		fmt.Println("mkfs succeeded")
	}
}
