package OssOverView

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	connect "yrcf_grpc_case/grpc/common"
	ossoverview "yrcf_grpc_case/grpc/ossoverview/proto"
)

func Run(client *ossoverview.OssOverviewClient, param *ossoverview.OssOverviewRequest) *ossoverview.OssOverviewRet {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := (*client).OssOverview(ctx, param)
	if err != nil {
		log.Fatal("did not get oss over view err: ", err)
	}

	if result.GetResult().GetErrorCode() != 0 {
		log.Printf("get oss over view failed, errcode=%d, result=%s\n",
			result.GetResult().GetErrorCode(), result.GetResult().GetResult())
		return nil
	} else {
		return result.GetValue()
	}
}

func Example(address string, port int) {
	fmt.Print("{\n  \"unused\": false,\n  \"cidr\": \"\"\n}\n")
	conn := connect.GetConn(address, port)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("did not close err:", err)
		}
	}(conn)

	client := ossoverview.NewOssOverviewClient(conn)
	results := Run(&client, &ossoverview.OssOverviewRequest{
		Unused: false,
		Cidr:   ""})

	printResult(results)
}

func printResult(results *ossoverview.OssOverviewRet) {
	fmt.Println("Oss Node Info:")
	nodeInfo := results.GetNodeInfo()
	var online = "offline"
	for _, node := range nodeInfo {
		if node.GetOnline() {
			online = "online"
		}

		fmt.Printf("[%d][%s][%s]\n", node.GetNodeNumId(), node.GetNodeName(), online)
	}

	fmt.Println("Space Used:")
	fmt.Printf("[%.2fMB/%.2fGB]\n", float64(results.GetDiskSpaceUsed())/1024/1024, float64(results.GetDiskSpaceTotal())/1024/1024/1024)
	fmt.Printf("Read/Write:\n%d/%d\n", results.GetDiskPerfRead()[0].GetValue(), results.GetDiskPerfWrite()[0].GetValue())
	fmt.Printf("Average Read/Write:\n%d/%d\n", results.GetDiskPerfAverageRead()[0].GetValue(), results.GetDiskPerAverageWrite()[0].GetValue())
	fmt.Printf("Disk Read/Write count:\n%d/%d\n", results.GetDiskReadSum(), results.GetDiskWriteSum())
	workRequests := results.GetWorkRequests()
	timestamp := workRequests[0].GetTime() / 10000
	count := uint64(0)
	fmt.Println("Work Requests:")
	for _, work := range workRequests {
		if timestamp == work.GetTime()/10000 {
			count += work.Value
		} else {
			fmt.Println(time.Unix(int64(timestamp*10), 0).Format("2006-01-02 15:04:05"), " count:", count)
			timestamp = work.GetTime() / 10000
			count = work.Value
		}
	}

	queuedRequests := results.GetQueuedRequests()
	timestamp = workRequests[0].GetTime() / 10000
	count = uint64(0)
	fmt.Println("Queue Requests:")
	for _, req := range queuedRequests {
		if timestamp == req.GetTime()/10000 {
			count += req.Value
		} else {
			fmt.Println(time.Unix(int64(timestamp*10), 0).Format("2006-01-02 15:04:05"), " count:", count)
			timestamp = req.GetTime() / 10000
			count = req.Value
		}
	}

}
