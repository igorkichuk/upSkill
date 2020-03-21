package servers

import (
	"context"
	"github.com/igorkichuk/upSkill/rpc/proto"
)

type CustomGreaterServer struct {

}

func (s *CustomGreaterServer) SayHello(ctxt context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message:              req.Name+" response",
	}, nil
}

func (s *CustomGreaterServer) SayHelloAgain(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello again " + in.GetName()}, nil
}