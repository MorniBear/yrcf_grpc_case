package MkDir

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	mkdir "yrcf_grpc_case/grpc/mkdir/proto"
)

func Run(client *mkdir.MkDirClient, param *mkdir.MkDirRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).MkDir(ctx, param)
	if err != nil {
		log.Fatal("did not mkdir, err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Mkdir failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return false
	} else {
		return true
	}
}

func Example(address string, port int) {
	println("{\n  " +
		"\"use_absolute_path\": false,\n  " +
		"\"path\": \"grpc_mkdir\",\n  " +
		"\"prefer_mds\": \"\",\n  " +
		"\"access_mode\": \"\",\n  " +
		"\"uid\": 0,\n  \"gid\": 0,\n  " +
		"\"not_use_mirror\": false,\n  " +
		"\"casefold\": false\n" +
		"}")
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := mkdir.NewMkDirClient(conn)
	result := Run(&client, &mkdir.MkDirRequest{
		UseAbsolutePath: false,
		Path:            "grpc_mkdir",
		PreferMds:       "",
		AccessMode:      "0755",
		Uid:             0,
		Gid:             0,
		NotUseMirror:    false,
		Casefold:        false,
	})
	if result {
		fmt.Println("mkdir succeeded")
	}
}
