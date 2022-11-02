package config

import "fmt"

func GetRoomHashKey(roomId string) string {
	return fmt.Sprintf("room-hash-key-%s", roomId)
}

func GetRoomMessageListKey(roomId string) string {
	return fmt.Sprintf("room-message-list-key-%s", roomId)
}

func GetWaitingRoomKey() string {
	return "waiting-room-key"
}
