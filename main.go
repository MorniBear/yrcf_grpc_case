package main

import (
	"yrcf_grpc_case/grpc/acl/request/DeleteAcl"
)

const (
	address = "10.16.13.121"
	port    = 8000
)

func main() {
	DeleteAcl.Example(address, port)
}
