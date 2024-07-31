package AddOrTestBucket

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	bucket "yrcf_grpc_case/grpc/bucket/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucket.BucketClient, param *bucket.AddOrTestBucketRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddOrTestBucket(ctx, param)
	if err != nil {
		log.Fatal("did not add or test bucket, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("add or test bucket failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucket.AddOrTestBucketRequest{
		Op: bucket.AddOrTestBucketRequest_ADD_BUCKET,
		BucketConfig: &bucket.BucketConfig{
			HostName:        "http://172.16.1.156:9000",
			BucketName:      "mybucket",
			Protocol:        1,
			UriStyle:        2,
			Region:          "",
			AccessKey:       "minioadmin",
			SecretAccessKey: "minioadmin",
			LibType:         bucket.BucketConfig_AWS,
			Token:           "",
		},
		AddWithoutTest: false,
	})

	if result {
		fmt.Println("add or test bucket succeeded")
	}
}
