package ShowBucketLinkHistory

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	bucketlink "yrcf_grpc_case/grpc/bucketlink/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *bucketlink.BucketLinkClient, param *bucketlink.ShowBucketLinkHistoryRequest) []*bucketlink.ExportHistoryInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ShowBucketLinkHistory(ctx, param)
	if err != nil {
		log.Fatal("did not show bucketlinks history, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Show bucketlinks history failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetExportHistoryInfo()
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
	result := Run(&client, &bucketlink.ShowBucketLinkHistoryRequest{
		LinkId:   1,
		ShowType: bucketlink.ShowBucketLinkHistoryRequest_EXPORT,
	})

	printResult(result)
}

func printResult(result []*bucketlink.ExportHistoryInfo) {

}
