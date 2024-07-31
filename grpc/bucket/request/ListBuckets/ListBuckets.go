package ListBuckets

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucket "yrcf_grpc_case/grpc/bucket/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucket.BucketClient, param *bucket.ListBucketsRequest) []*bucket.BucketInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListBuckets(ctx, param)
	if err != nil {
		log.Fatal("did not list buckets err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("list buckets failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetBucketInfo()
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
	results := Run(&client, &bucket.ListBucketsRequest{})

	printResult(results)
}

func printResult(results []*bucket.BucketInfo) {}
