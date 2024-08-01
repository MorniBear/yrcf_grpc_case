package GetLicense

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	license "yrcf_grpc_case/grpc/license/proto"
)

func Run(client *license.LicenseClient, param *license.GetLicenseRequest) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).GetLicense(ctx, param)
	if err != nil {
		log.Fatal("did not get license, err: ", err)
	}

	return result.GetTime()
}

func Example(address string, port int) {
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := license.NewLicenseClient(conn)
	result := Run(&client, &license.GetLicenseRequest{})

	if result != "" {
		fmt.Println(result)
	}
}
