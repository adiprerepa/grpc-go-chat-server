package chat
import (
	"context"
	chat "go_generated"
)

type chatServiceServer struct {}

func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{}
}

func (c chatServiceServer) RegisterRoom(ctx context.Context, message *chat.RegisterClient) (*chat.RegisterClientResponse, error) {
    //panic("implement me")

    enumVal, ok := chat.RegisterClientResponseAuthResponseStatus_value["OK"]
    if !ok {
    	panic("invalid enum thing")
	}
	return &chat.RegisterClientResponse{
		Status:               chat.RegisterClientResponseAuthResponseStatus.Enum(enumVal),
		Users:                nil,
	}, nil
    // return 
}

func (c chatServiceServer) SendMessage(context.Context, *chat.ChatMessage) (*chat.ChatMessageResponse, error) {
	panic("implement me")
}

func (c chatServiceServer) UpdateMessage(*chat.MessageUpdateReq, chat.ChatService_UpdateMessageServer) error {
	panic("implement me")
}

func (c chatServiceServer) ExitChat(context.Context, *chat.ExitRequest) (*chat.ExitResponse, error) {
	panic("implement me")
}
