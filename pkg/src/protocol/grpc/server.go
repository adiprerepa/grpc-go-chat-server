package grpc

import (
	"context"
	api "go_generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
  Error is for when it is already binded?
*/
func RunServer(ctx context.Context, srv api.ChatServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":" + port)

	if err != nil {
		return err
	}

	// register service with netty
	server := grpc.NewServer()
	api.RegisterChatServiceServer(server, srv)

	// start grpc
	log.Printf("Starting gRPC server on port %v....\n", port)
	return server.Serve(listen)
}
