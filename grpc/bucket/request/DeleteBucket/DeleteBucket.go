package DeleteBucket

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	bucket "yrcf_grpc_case/grpc/bucket/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucket.BucketClient, param *bucket.DeleteBucketRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteBucket(ctx, param)
	if err != nil {
		log.Fatal("did not delete bucket, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("delete bucket failed, errcode=%d, result=%s\n",
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

	client := bucket.NewBucketClient(conn)
	result := Run(&client, &bucket.DeleteBucketRequest{
		BucketId: 1,
	})

	if result {
		fmt.Println("delete bucket succeeded")
	}
}
