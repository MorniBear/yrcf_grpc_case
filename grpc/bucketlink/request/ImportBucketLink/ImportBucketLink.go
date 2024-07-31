package ImportBucketLink

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.ImportBucketLinkRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ImportBucketLink(ctx, param)
	if err != nil {
		log.Fatal("did not import bucketlinks, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Import bucketlinks failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucketlink.ImportBucketLinkRequest{
		LinkId: 1,
		ImportConfig: &bucketlink.ImportConfig{
			Scope:     bucketlink.ImportConfig_IMPORT_ALL,
			Pattern:   "",
			Prefix:    "",
			LoadType:  bucketlink.ImportConfig_LOAD_META,
			Threads:   2,
			ImportAll: true,
		},
		ImportLater: false,
	})

	if result {
		log.Println("success")
	}
}
