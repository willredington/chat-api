package service

import (
	"com/willredington/chat-api/config"
	"com/willredington/chat-api/model"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func CreateRoom(client *redis.Client, userIds []string) (*model.Room, error) {

	id := uuid.New()

	room := &model.Room{
		Id:           id.String(),
		UserIds:      userIds,
		Status:       model.ROOM_STATUS_IDLE,
		CreationDate: time.Now().String(),
	}

	key := config.GetRoomHashKey(room.Id)

	roomJson, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	pipeline := client.TxPipeline()

	pipeline.HSet(client.Context(), key, roomJson)

	for _, userId := range userIds {
		pipeline.Set(client.Context(), userId, room.Id, 0)
	}

	pipeline.Exec(client.Context())

	return room, nil
}

func FindRoomForUser(client *redis.Client, userId string) (string, error) {
	return client.Get(client.Context(), userId).Result()
}

func SendMessage(client *redis.Client, roomId string, message *model.Message) error {

	msgJson, err := json.Marshal(message)
	if err != nil {
		return nil
	}

	return client.LPush(client.Context(), config.GetRoomMessageListKey(roomId), msgJson).Err()
}

func GetMessages(client *redis.Client, roomId string) ([]*model.Message, error) {

	vals, err := client.LRange(client.Context(), config.GetRoomMessageListKey(roomId), -50, -1).Result()
	if err != nil {
		return nil, err
	}

	var result []*model.Message

	for _, val := range vals {
		var message = &model.Message{}
		if err := json.Unmarshal([]byte(val), &message); err != nil {
			return nil, err
		}

		result = append(result, message)
	}

	return result, nil
}

func PushToWaitQueue(client *redis.Client, userIds []string) error {
	return client.SAdd(client.Context(), config.GetWaitingRoomKey(), userIds).Err()
}

func IsInWaitQueue(client *redis.Client, userId string) (bool, error) {
	return client.SIsMember(client.Context(), config.GetWaitingRoomKey(), userId).Result()
}

func PopWaitQueue(client *redis.Client) (string, error) {
	return client.SPop(client.Context(), config.GetWaitingRoomKey()).Result()
}
