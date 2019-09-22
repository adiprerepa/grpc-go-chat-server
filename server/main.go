package main

import (
	"context"
	"fmt"
	"os"
	"protocol/grpc"
	"service/chat"
)

func main() {
	if err := grpc.RunServer(context.Background(), chat.NewChatServiceServer(), "3000"); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
