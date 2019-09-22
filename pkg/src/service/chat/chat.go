package chat
import (
	"context"
	chat "go_generated"
	modelstore "modelstore/value"
)

type chatServiceServer struct {}

func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{}
}

// add yourself to a room + get all users
func (c chatServiceServer) RegisterRoom(ctx context.Context, message *chat.RegisterClient) (*chat.RegisterClientResponse, error) {
    //panic("implement me")
    // get rooms and status of rooms
    stat, rooms := modelstore.AddUser(*message.RoomId, *message.Username)
    var status int32
    // if there is a room
    if stat == true {
    	status = 1
    	// all users + stat
    	return &chat.RegisterClientResponse{
			Status:  &status,
			Users:  rooms,
		}, nil
	} else {
		status = 2
		return &chat.RegisterClientResponse{
			Status: &status,
			Users: nil,
		}, nil
	}
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
