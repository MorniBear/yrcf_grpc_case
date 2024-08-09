package SetEntry

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	setentry "yrcf_grpc_case/grpc/setentry/proto"
)

func Run(client *setentry.SetEntryClient, param *setentry.SetEntryRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).SetEntry(ctx, param)
	if err != nil {
		log.Fatal("did not set entry, err: ", err)
	}
	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Set entry failed, errcode=%d, result=%s\n",
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

	client := setentry.NewSetEntryClient(conn)
	result := Run(&client, &setentry.SetEntryRequest{
		UseMountedPath: false,
		Path:           "/",
		StripeSize:     "1m",
		StripeCount:    3,
		Schema:         setentry.SetEntryRequest_EC,
		PoolId:         "",
		ForceSet:       false,
		EcBlocksize:    "1m",
	})

	if result {
		fmt.Println("set entry succeeded")
	}
}
