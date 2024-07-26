package ListAcl

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	acl "yrcf_grpc_case/grpc/acl/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *acl.AclClient, param *acl.ListAclRequest) []*acl.AclInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListAcl(ctx, param)
	if err != nil {
		log.Fatal("did not list acl, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("list acl failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetAclInfoList()
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
	result := Run(&client, &acl.ListAclRequest{
		Path: "/",
	})

	printResult(result)
}

func printResult(result []*acl.AclInfo) {
	for _, item := range result {
		println("path:", item.GetPath())
		println("resp:", item.GetResponse())
	}
}
