package GetQosSla

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	getqossla "yrcf_grpc_case/grpc/getqossla/proto"
)

func Run(client *getqossla.GetQosSlaClient, param *getqossla.GetQosSlaRequest) []*getqossla.QosSla {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).GetQosSla(ctx, param)
	if err != nil {
		log.Fatal("did not get qos sla err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Get qos sla failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetQosSla()
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

	client := getqossla.NewGetQosSlaClient(conn)
	result := Run(&client, &getqossla.GetQosSlaRequest{
		UseAbsolutePath: false,
		Path:            "/grpc_qos",
	})

	printResult(result)
}

func printResult([]*getqossla.QosSla) {

}
