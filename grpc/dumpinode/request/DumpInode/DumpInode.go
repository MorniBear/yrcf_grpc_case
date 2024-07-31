package DumpInode

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	dumpinode "yrcf_grpc_case/grpc/dumpinode/proto"
)

func Run(client *dumpinode.DumpInodeClient, param *dumpinode.DumpInodeRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).DumpInode(ctx, param)
	if err != nil {
		log.Fatal("did not dump inode, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Dump inode failed, errcode=%d, result=%s\n",
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

	client := dumpinode.NewDumpInodeClient(conn)
	result := Run(&client, &dumpinode.DumpInodeRequest{
		UseAbsolutePath: false,
		Path:            "/",
		Type:            dumpinode.DumpInodeRequest_MDS,
		EntryId:         "1006566A77BF9",
	})
	if result {
		log.Println("success")
	}
}
