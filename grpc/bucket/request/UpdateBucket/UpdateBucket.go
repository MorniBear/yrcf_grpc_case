package UpdateBucket

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	bucket "yrcf_grpc_case/grpc/bucket/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucket.BucketClient, param *bucket.UpdateBucketRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).UpdateBucket(ctx, param)
	if err != nil {
		log.Fatal("did not update bucket, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("update bucket failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucket.UpdateBucketRequest{
		BucketId:        1,
		AccessKey:       "minioadmin",
		SecretAccessKey: "minioadmin",
	})

	if result {
		fmt.Println("update bucket succeeded")
	}
}
