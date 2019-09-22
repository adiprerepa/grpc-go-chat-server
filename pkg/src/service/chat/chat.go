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
func (c chatServiceServer) RegisterClient(ctx context.Context, message *chat.RegisterClient) (*chat.RegisterClientResponse, error) {
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

func (c chatServiceServer) SendMessage(ctx context.Context, message *chat.ChatMessage) (*chat.ChatMessageResponse, error) {
	broadcastStatus := modelstore.BroadcastMsg(*message.Username, *message.Payload, *message.RoomId)
	var status int32
	if broadcastStatus {
		status = 1
		return &chat.ChatMessageResponse{
			MesssageStatus:  &status,
		}, nil
	} else {
		status = 2
		return &chat.ChatMessageResponse{
			MesssageStatus:  &status,
		}, nil
	}
}

func (c chatServiceServer) UpdateMessage(req *chat.MessageUpdateReq, observable chat.ChatService_UpdateMessageServer) error {
	// called when broadcast in server is called
	stat := modelstore.AddObservable(*req.ChatRoomId, *req.Username, observable)
	if stat {
		return nil
	} else {
		panic("err with observable addition")
	}
}

func (c chatServiceServer) ExitChat(ctx context.Context, exitReq *chat.ExitRequest) (*chat.ExitResponse, error) {
	var stat int32
	if modelstore.RemoveUser(*exitReq.Username, *exitReq.ChatId) {
		stat = 1
		return &chat.ExitResponse{
			ExitStatus: &stat,
		}, nil
	} else {
		stat = 2
		return &chat.ExitResponse{
			ExitStatus: &stat,
		}, nil
	}
}
