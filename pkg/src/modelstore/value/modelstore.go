package value

import chat "go_generated"

type user struct {
	username string
	// todo add cur chat room?
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
			nrmToPb(room)
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