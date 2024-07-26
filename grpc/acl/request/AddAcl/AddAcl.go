package AddAcl

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	acl "yrcf_grpc_case/grpc/acl/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *acl.AclClient, param *acl.AddAclRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddAcl(ctx, param)
	if err != nil {
		log.Fatal("did not add acl, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("add acl failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &acl.AddAclRequest{
		Path: "/",
		Ip:   "192.168.0.1",
		Id:   "",
		Mode: acl.AddAclRequest_RW,
	})

	if result {
		fmt.Println("add acl succeeded")
	}
}
