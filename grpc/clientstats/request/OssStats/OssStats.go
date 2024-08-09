package OssStats

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	"yrcf_grpc_case/grpc/clientstats/proto"
	"yrcf_grpc_case/grpc/common"
)

func Run(client *clientstats.ClientStatsClient, param *clientstats.ClientStatsRequest) []*clientstats.ClientStatsRet {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).OssStats(ctx, param)
	if err != nil {
		log.Fatal("did not get oss stat err:", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("get oss stats failed, errcode=%d, result=%s\n", result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetValue()
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

	client := clientstats.NewClientStatsClient(conn)
	results := Run(&client, &clientstats.ClientStatsRequest{
		HideInternalIps: true,
		ReturnZeroStats: true,
		ClientStatsType: clientstats.ClientStatsRequest_EXT,
		Cidr:            ""})

	printResult(results)
}

func printResult(results []*clientstats.ClientStatsRet) {
	linen := 5
	ossOpStrSet := []string{
		"sum",
		"ack",
		"s_ch_drct",
		"get_file_size",
		"s_attr",
		"statfs",
		"trunc",
		"close",
		"fsync",
		"open",
		"iops_rd",
		"bps_rd",
		"iops_wr",
		"bps_wr",
		"gendbg",
		"hrtbeat",
		"rem_node",
		"node_inf",
		"stor_info",
		"unlnk",
		"reserve",
		"reserve",
		"reserve",
		"reserve"}

	for i, result := range results {
		fmt.Printf("[%d][%s]", i, result.Ip)
		if result.Ip != "total" {
			if result.Online {
				fmt.Println("[online]")
			} else {
				fmt.Println("[offline]")
			}
		} else {
			fmt.Println()
		}

		fmt.Println("------------------------------------------------------------------------------------------")
		for i := 0; i < len(ossOpStrSet); i += linen {
			for j := i; j < i+linen && j < len(ossOpStrSet); j++ {
				fmt.Printf("%-20s", ossOpStrSet[j])
			}
			fmt.Println()
			for j := i; j < i+linen && j < len(ossOpStrSet); j++ {
				fmt.Printf("%-20d", result.Opcounters[j])
			}
			fmt.Println("")
		}
		fmt.Println("------------------------------------------------------------------------------------------")
		fmt.Println()
	}
}
