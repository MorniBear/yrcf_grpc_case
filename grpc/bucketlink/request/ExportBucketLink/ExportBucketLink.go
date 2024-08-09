package ExportBucketLink

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.ExportBucketLinkRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ExportBucketLink(ctx, param)
	if err != nil {
		log.Fatal("did not Export bucketlinks, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Export bucketlinks failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucketlink.ExportBucketLinkRequest{
		LinkId: 1,
		ExportConfig: &bucketlink.ExportConfig{
			BucketId:    1,
			Scope:       bucketlink.ExportConfig_EXPORT_ALL,
			Pattern:     "",
			Prefix:      "",
			NameRule:    bucketlink.ExportConfig_USE_ORIGIN,
			NameSuffix:  "",
			WillPurge:   true,
			PurgeTiming: bucketlink.ExportConfig_AFTER_EXPORT,
			Threads:     2,
			EtagScope:   bucketlink.ExportConfig_DONT_CHECK_ETAG,
		},
		ExportLater: false,
	})

	if result {
		log.Println("success")
	}
}
