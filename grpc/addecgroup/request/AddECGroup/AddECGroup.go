package AddECGroup

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	addecgroup "yrcf_grpc_case/grpc/addecgroup/proto"
	connect "yrcf_grpc_case/grpc/common"
)

func Run(client *addecgroup.AddECGroupClient, param *addecgroup.AddECGroupRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).AddECGroup(ctx, param)
	if err != nil {
		log.Fatal("did not add EC group, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("add ec group failed, errcode=%d, result=%s\n",
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

	client := addecgroup.NewAddECGroupClient(conn)
	var result = Run(&client, &addecgroup.AddECGroupRequest{
		AutoAddGroup: false,
		ForceAdd:     true,
		Preview:      false,
		EcGroupInfo: &addecgroup.ECGroupInfo{
			GroupId:     1,
			DataNum:     4,
			ParityNum:   2,
			MasterOsdId: 0,
			OsdIds:      "",
		},
	})
	if result {
		fmt.Println("add ec group succeeded")
	}
}
