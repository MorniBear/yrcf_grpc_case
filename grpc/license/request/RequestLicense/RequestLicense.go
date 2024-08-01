package RequestLicense

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	license "yrcf_grpc_case/grpc/license/proto"
)

func Run(client *license.LicenseClient, param *license.RequestLicenseRequest) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).RequestLicense(ctx, param)
	if err != nil {
		log.Fatal("did not request license, err: ", err)
	}

	return result.String()
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
	result := Run(&client, &license.RequestLicenseRequest{})

	if result != "" {
		fmt.Println(result)
	}
}
