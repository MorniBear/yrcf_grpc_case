package GetQuotaSla

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	getquotasla "yrcf_grpc_case/grpc/getquotasla/proto"
)

func Run(client *getquotasla.GetQuotaSlaClient, param *getquotasla.GetQuotaSlaRequest) []*getquotasla.QuotaSla {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).QuotaSla(ctx, param)
	if err != nil {
		log.Fatal("did not get quota sla err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("Get quota sla failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetQuotaSla()
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

	client := getquotasla.NewGetQuotaSlaClient(conn)
	result := Run(&client, &getquotasla.GetQuotaSlaRequest{
		WithRoot: false,
		Path:     "/grpc_projectquota",
	})

	printResult(result)
}

func printResult(result []*getquotasla.QuotaSla) {

}
