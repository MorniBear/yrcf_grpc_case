package SetLicense

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	license "yrcf_grpc_case/grpc/license/proto"
)

func Run(client *license.LicenseClient, param *license.SetLicenseRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).SetLicense(ctx, param)
	if err != nil {
		log.Fatal("did not set license, err: ", err)
	}

	log.Println(result.GetResult().GetResult())
}

func Example(address string, port int) {
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)
	path := "/root/yrcloud_file_license.yr"
	client := license.NewLicenseClient(conn)
	Run(&client, &license.SetLicenseRequest{
		Path: &path,
	})
}
