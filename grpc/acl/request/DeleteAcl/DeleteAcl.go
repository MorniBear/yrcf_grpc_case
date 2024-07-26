package DeleteAcl

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	acl "yrcf_grpc_case/grpc/acl/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *acl.AclClient, param *acl.DeleteAclRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteAcl(ctx, param)
	if err != nil {
		log.Fatal("did not del acl, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("del acl failed, errcode=%d, result=%s\n",
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

	client := acl.NewAclClient(conn)
	result := Run(&client, &acl.DeleteAclRequest{
		Path: "/",
		Ip:   "192.168.1.1",
	})

	if result {
		fmt.Println("del acl succeeded")
	}
}
