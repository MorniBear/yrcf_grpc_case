package BucketLinkSubscribeOps

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.BucketLinkSubscribeOpsRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).BucketLinkSubscribeOps(ctx, param)
	if err != nil {
		log.Fatal("did not subscribe bucketlinks, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Subscribe bucketlinks failed, errcode=%d, result=%s\n",
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

	client := bucketlink.NewBucketLinkClient(conn)
	result := Run(&client, &bucketlink.BucketLinkSubscribeOpsRequest{
		LinkId: 1,
		SubOp:  0,
		SubConfig: &bucketlink.SubscribeConfig{
			S3Type:         bucketlink.SubscribeConfig_YANRONGYUN,
			Prefix:         "",
			SubWithoutTest: false,
		},
	})

	if result {
		log.Println("success")
	}
}
