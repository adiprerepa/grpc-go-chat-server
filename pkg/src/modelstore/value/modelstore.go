package value

import (
	chat "go_generated"
	"time"
)

type user struct {
	username string
	observable chat.ChatService_UpdateMessageServer
}

type chatRoom struct {
	users[] user
	chatRoomId string
}

var chatRooms[] chatRoom


// test for case when roomId doesnt exist
func AddUser(roomId string, username string) (bool, []*chat.RoomUserEntity) {
	var usersInChat []* chat.RoomUserEntity
	roomExists := false
	roomUser := user{username:username}
	for _, room := range chatRooms {
		if room.chatRoomId == roomId {
			roomExists = true
			usersInChat = nrmToPb(room)
			room.users = append(room.users, roomUser)
		}
	}
	return roomExists, usersInChat
}

func nrmToPb(chatRoom chatRoom) (pbRoom []*chat.RoomUserEntity) {
	users := chatRoom.users
	var pbRooms[]* chat.RoomUserEntity
	for _, user := range users {
		pbRooms = append(pbRooms,
			&chat.RoomUserEntity{
				Username: &user.username})
	}
	return pbRooms
}

// maybe add return?
func BroadcastMsg(username, payload, chatId string) bool {
	var ret = true
	for _, chatRoom := range chatRooms {
		if chatId == chatRoom.chatRoomId {
			for _, user := range chatRoom.users {
				strTime := time.Now().String()
				msg := chat.ChatMessage{
					Username:             &username,
					Payload:              &payload,
					Time:                 &strTime,
				}
				updateMsg := chat.UpdateChatMessage{
					Message: &msg,
				}
				err := user.observable.Send(&updateMsg)
				if err != nil {
					ret = false
					panic(err)
				}
			}
		}
	}
	return ret
}

func AddObservable(roomId string, username string, observable chat.ChatService_UpdateMessageServer) {
	var found = false
	for _, chatRoom := range chatRooms {
		// find right room
		if chatRoom.chatRoomId == roomId {
			found = true
			for _, user := range chatRoom.users {
				// find right user to add observable to
				if username == user.username {
					user.observable = observable
				}
			}
		}
	}
	// create a new room
	if found == false {
		chatUser := user{
			username:   username,
			observable: observable,
		}
		var chatRoomUsers []user
		room := chatRoom{
			// new arr with 1 user
			users: append(chatRoomUsers, chatUser),
			chatRoomId: roomId,
		}
		// append to chatrooms
		chatRooms = append(chatRooms, room)
	}
}

func RemoveUser(roomId, username string) bool {
	var ret = false
	for _, chatRoom := range chatRooms {
		if chatRoom.chatRoomId == roomId {
			numOfUsers := 0
			for _, user := range chatRoom.users {
				if user.username == username {
					// remove
					chatRoom.users = append(chatRoom.users[:numOfUsers], chatRoom.users[numOfUsers+1:]...)
					ret = true
				}
				numOfUsers++
			}
		}
	}
	return ret
}