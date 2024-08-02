package SetAttr

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	setattr "yrcf_grpc_case/grpc/setattr/proto"
)

func Run(client *setattr.SetAttrClient, param *setattr.SetAttrRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).SetAttr(ctx, param)
	if err != nil {
		log.Fatal("did not set attr err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Set attr failed, errcode=%d, result=%s\n",
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

	mTimeSecs := int64(1000)
	aTimeSecs := int64(1000)
	client := setattr.NewSetAttrClient(conn)
	result := Run(&client, &setattr.SetAttrRequest{
		UseAbsolutePath: false,
		Path:            "/",
		AccessMode:      "0777",
		Uid:             0,
		Gid:             0,
		MtimeSecs:       &mTimeSecs,
		MtimeNano:       nil,
		AtimeSecs:       &aTimeSecs,
		AtimeNano:       nil,
	})
	if result {
		fmt.Println("setattr succeeded")
	}
}
