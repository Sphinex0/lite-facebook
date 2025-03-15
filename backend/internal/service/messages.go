package service

import (
	"fmt"
	"time"

	"social-network/internal/models"
)

func (service *Service) CreateMessage(msg *models.Message) (err error) {
	msg.CreatedAt = int(time.Now().Unix())
	err = service.Database.SaveMessage(msg)
	return
}

func (service *Service) FetchMessagesHestories(befor, conversation_id int) (messages []models.WSMessage, err error) {
	fmt.Println(befor, conversation_id)
	messages, err = service.Database.GetMessagesHestories(befor, conversation_id)
	return
}
