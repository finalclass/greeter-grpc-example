package main

import (
	"log"
	"net"

	"greeter/pkg/pb"
	"greeter/pkg/service"

	"google.golang.org/grpc"
)

func main() {
	port := "8080"
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &service.Server{})

	log.Printf("Greeter listening on 127.0.0.1:" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
