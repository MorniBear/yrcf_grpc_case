package SetCidrPriority

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	cidrpriority "yrcf_grpc_case/grpc/cidrpriority/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *cidrpriority.CidrPriorityClient, param *cidrpriority.SetCidrPriRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).SetCidrPriority(ctx, param)
	if err != nil {
		log.Fatal("did not get cidr priority, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Get cidr priority failed, errcode=%d, result=%s\n",
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

	client := cidrpriority.NewCidrPriorityClient(conn)
	result := Run(&client, &cidrpriority.SetCidrPriRequest{
		Value: "net1.net2",
	})
	if result {
		log.Println("success")
	}
}
