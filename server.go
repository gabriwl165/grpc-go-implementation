package main

import (
	"log"
	"net"

	"github.com/gabriwl165/grpc-go-implementation/chat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}

}
