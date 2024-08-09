package MdsStats

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
	result, err := (*client).MdsStats(ctx, param)
	if err != nil {
		log.Fatal("did not get mds stat err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("get mds stats failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
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
		HideInternalIps: false,
		ReturnZeroStats: true,
		ClientStatsType: clientstats.ClientStatsRequest_EXT,
		Cidr:            ""})

	printResult(results)
}

func printResult(results []*clientstats.ClientStatsRet) {
	linen := 8
	mdsOpStrSet := []string{
		"sum",
		"ack",
		"close",
		"ent_inf",
		"node_inf",
		"fnd_own",
		"lnk",
		"mkdir",
		"create",
		"rddir",
		"refr_ent",
		"mds_inr",
		"rmlnk",
		"mv_dir_ins",
		"mv_file_ins",
		"open",
		"ren",
		"s_ch_drct",
		"s_attr",
		"s_dir_pat",
		"stat",
		"statfs",
		"trunc",
		"symlnk",
		"unlnk",
		"look_li",
		"stat_li",
		"reval_li",
		"open_li",
		"create_li",
		"mirror_md",
		"hardlnk",
		"flck_ap",
		"flck_en",
		"flck_rg",
		"dirparent",
		"list_xa",
		"get_xa",
		"rm_xa",
		"set_x"}

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
		for i := 0; i < len(mdsOpStrSet); i += linen {
			for j := i; j < i+linen && j < len(mdsOpStrSet); j++ {
				fmt.Printf("%-12s", mdsOpStrSet[j])
			}
			fmt.Println()
			for j := i; j < i+linen && j < len(mdsOpStrSet); j++ {
				fmt.Printf("%-12d", result.Opcounters[j])
			}
			fmt.Println("")
		}
		fmt.Println("------------------------------------------------------------------------------------------")
		fmt.Println()
	}
}
