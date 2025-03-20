package service

import (
	"time"

	"social-network/internal/models"
)

func (service *Service) CreateMessage(msg *models.WSMessage) (err error) {
	msg.Message.CreatedAt = int(time.Now().UnixMilli())
	err = service.Database.SaveMessage(msg)
	return
}

func (service *Service) FetchMessagesHestories(befor, conversation_id int) (messages []models.WSMessage, err models.Error) {
	messages, err = service.Database.GetMessagesHestories(befor, conversation_id)
	return
}

func (service *Service) ReadMessages(convId int) (err error) {
	err = service.Database.ReadMessages(convId)
	return
}
