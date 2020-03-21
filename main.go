package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/igorkichuk/upSkill/pb"
)

const port = ":1280"

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	_, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
