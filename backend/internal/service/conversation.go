package service

import (
	"social-network/internal/models"
)

func (service *Service) FetchConversations(id int) (conversations []models.ConversationsInfo, err error) {
	service.Database.GetConversations(id)
	return
}

func (service *Service) CreateConversation(conv *models.Conversation) (conversations []models.ConversationsInfo, err error) {
	service.Database.CreateConversation(conv)
	return
}
