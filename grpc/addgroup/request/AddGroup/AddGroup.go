package AddGroup

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	addgroup "yrcf_grpc_case/grpc/addgroup/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *addgroup.AddGroupClient, param *addgroup.AddGroupRequest) []*addgroup.OsdGroupInfo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddGroup(ctx, param)
	if err != nil {
		log.Fatal("did not add group, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("add acl group, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetOsdGroupInfo()
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

	client := addgroup.NewAddGroupClient(conn)
	result := Run(&client, &addgroup.AddGroupRequest{
		OsdType:       addgroup.AddGroupRequest_MDS,
		AutoAddGroup:  true,
		MasterNode:    0,
		SlaveNode:     0,
		GroupId:       0,
		ForceAdd:      false,
		Preview:       false,
		UniqueGroupId: false,
	})

	printResult(result)
}

func printResult(result []*addgroup.OsdGroupInfo) {

}
