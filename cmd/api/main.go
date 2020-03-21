package main

import (
	"fmt"
	"google.golang.org/grpc"
	"mod/rpc/proto"
	"mod/servers"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 1280))
	if err != nil {
		fmt.Println(err)
		return
	}

	grpcServer := grpc.NewServer()
	grServer := new(servers.CustomGreaterServer)
	proto.RegisterGreeterServer(grpcServer, grServer)
	grpcServer.Serve(lis)
}