package main

import (
	"context"
	"fmt"
	"os"
	"protocol/grpc"
	"service/chat"
)

func main() {
	var port string
	if len(os.Args) >= 2 {
		port = os.Args[1]
	} else {
		port = "3000"
	}
	if err := grpc.RunServer(context.Background(), chat.NewChatServiceServer(), port); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
