package main

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
func addUser(roomId string, chatuser chat.RoomUserEntity) bool {
	roomExists := false
	roomUser := user{username:*chatuser.Username}
	for _, room := range chatRooms {
		if room.chatRoomId == roomId {
			roomExists = true
			append(room.users, roomUser)
		}
	}
	return roomExists
}