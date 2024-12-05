package service

import (
	"context"
	"greeter/pkg/pb"

	"google.golang.org/protobuf/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := req.GetName()

	return &pb.HelloReply{
		Message: *proto.String("Hello " + name),
	}, nil

}
