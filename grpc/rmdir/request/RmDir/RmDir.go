package RmDir

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	rmdir "yrcf_grpc_case/grpc/rmdir/proto"
)

func Run(client *rmdir.RmDirClient, param *rmdir.RmDirRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).RmDir(ctx, param)
	if err != nil {
		log.Fatal("did not rmdir err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Rmdir failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return false
	} else {

		return true
	}
}

func Example(address string, port int) {
	const path = "rm_dir"
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := rmdir.NewRmDirClient(conn)
	result := Run(&client, &rmdir.RmDirRequest{
		UseAbsolutePath:   false,
		RemoveRecursively: false,
		Path:              []string{path},
	})

	if result {
		fmt.Println("rmdir succeeded")
	}
}
