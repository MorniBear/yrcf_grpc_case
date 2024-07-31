package StatBucketLink

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.StatBucketLinkRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).StatBucketLink(ctx, param)
	if err != nil {
		log.Fatal("did not stat bucketlinks, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Stat bucketlinks failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucketlink.StatBucketLinkRequest{
		LinkId: 1,
	})
	if result {
		log.Println("success")
	}
}
