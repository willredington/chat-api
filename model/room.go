package model

const (
	ROOM_STATUS_IDLE   = "IDLE"
	ROOM_STATUS_ACTIVE = "ACTIVE"
	ROOM_STATUS_DEAD   = "DEAD"
)

type Room struct {
	Id           string   `json:"id"`
	UserIds      []string `json:"userIds"`
	Status       string   `json:"status"`
	CreationDate string   `json:"creationDate"`
}
