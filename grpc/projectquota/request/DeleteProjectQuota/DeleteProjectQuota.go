package DeleteProjectQuota

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	mkdir "yrcf_grpc_case/grpc/mkdir/proto"
	MkDir "yrcf_grpc_case/grpc/mkdir/request"
	projectquota "yrcf_grpc_case/grpc/projectquota/proto"
	"yrcf_grpc_case/grpc/projectquota/request/AddOrUpdateProjectQuota"
)

func Run(client *projectquota.ProjectQuotaClient, param *projectquota.DeleteProjectQuotaRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DeleteProjectQuota(ctx, param)
	if err != nil {
		log.Fatal("Del project quota err:", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Del project quota failed, errcode=%d, result=%s\n", result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return false
	}
	return true
}

func Example(address string, port int) {
	const path = "grpc_project_quota"
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	//first, we need to create a dir for project quota
	mkDirClient := mkdir.NewMkDirClient(conn)
	result := MkDir.Run(&mkDirClient, &mkdir.MkDirRequest{
		UseAbsolutePath: false,
		Path:            path,
		PreferMds:       "",
		AccessMode:      "0755",
		Uid:             0,
		Gid:             0,
		NotUseMirror:    false,
		Casefold:        false,
	})

	if !result {
		return
	}

	//second, create quota for this dir
	quotaClient := projectquota.NewProjectQuotaClient(conn)
	result = AddOrUpdateProjectQuota.Run(&quotaClient, &projectquota.AddOrUpdateProjectQuotaRequest{
		Op:              projectquota.AddOrUpdateProjectQuotaRequest_ADD,
		Path:            path,
		SpaceLimit:      "10G",
		InodeLimit:      1000,
		Recursive:       true,
		IgnoreExisting:  false,
		ProjectContinue: false,
	})

	if !result {
		return
	}
	//third, delete this quota
	result = Run(&quotaClient, &projectquota.DeleteProjectQuotaRequest{
		Path:  path,
		Eid:   "",
		Force: true,
	})

	if result {
		fmt.Println("Del project quota success")
	}
}
