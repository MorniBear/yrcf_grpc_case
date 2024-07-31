package HotLayerCleanupOps

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.HotLayerCleanupOpsRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).HotLayerCleanupOps(ctx, param)
	if err != nil {
		log.Fatal("did not clean up hot layer, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Clean up hot layer failed, errcode=%d, result=%s\n",
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
	result := Run(&client, &bucketlink.HotLayerCleanupOpsRequest{
		CleanupOp:    bucketlink.HotLayerCleanupOpsRequest_UPDATE_CLEANUP_RULES,
		LinkId:       1,
		ColdTime:     2,
		CleanupTimer: "",
	})

	if result {
		log.Println("success")
	}
}
