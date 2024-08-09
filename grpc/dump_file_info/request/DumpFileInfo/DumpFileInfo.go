package DumpFileInfo

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	dump_file_info "yrcf_grpc_case/grpc/dump_file_info/proto"
)

func Run(client *dump_file_info.DumpFileInfoClient, param *dump_file_info.DumpFileInfoRequest) *dump_file_info.DumpFileInfoResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DumpFileInfo(ctx, param)
	if err != nil {
		log.Fatal("did not dump file info, err: ", err)
	}

	if result.GetResult() != 0 {
		log.Printf("Dump file info failed, errcode=%d msg=%s\n",
			result.GetResult(), result.GetMsg())
		return nil
	} else {
		return result
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

	client := dump_file_info.NewDumpFileInfoClient(conn)
	result := Run(&client, &dump_file_info.DumpFileInfoRequest{
		Path: "/",
	})

	printResult(result)
}
func printResult(result *dump_file_info.DumpFileInfoResponse) {}
