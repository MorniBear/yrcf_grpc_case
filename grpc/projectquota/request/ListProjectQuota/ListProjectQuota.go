package ListProjectQuota

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	projectquota "yrcf_grpc_case/grpc/projectquota/proto"
)

func Run(client *projectquota.ProjectQuotaClient, param *projectquota.ListProjectQuotaRequest) []*projectquota.ProjectQuotaInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).ListProjectQuota(ctx, param)
	if err != nil {
		log.Fatal("List project quota err:", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("List project quota failed, errcode=%d, result=%s\n", result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	}
	return result.GetProjectQuotaInfoList()
}

func Example(address string, port int) {
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := projectquota.NewProjectQuotaClient(conn)
	result := Run(&client, &projectquota.ListProjectQuotaRequest{
		Path:   "",
		Prefix: false,
	})

	if result != nil {
		PrintResult(result)
	}
}

func PrintResult(results []*projectquota.ProjectQuotaInfo) {
	for _, info := range results {
		fmt.Printf("Quota_id: %d\nEntry_id: %s\nPath: %s\n"+
			"Space [%d/%d]\nInode [%d/%d]\n"+"status: %s\nrecursive: %t\n"+"dir_used: %d\nfile_used: %d\n"+
			"--------------------------\n",
			info.GetQuotaId(), info.GetEid(), info.GetPath(), info.GetSpaceUsed(),
			info.GetSpaceLimit(), info.InodeUsed, info.InodeLimit, info.GetOpStatus(),
			info.GetRecursive(), info.GetDirUsed(), info.GetFileUsed())
	}
}
