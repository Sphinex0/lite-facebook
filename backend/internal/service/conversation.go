package service

import (
	"social-network/internal/models"
)

func (service *Service) FetchConversations(id int) (conversations []models.ConversationsInfo, err error) {
	conversations, err = service.Database.GetConversations(id)
	return
}

func (service *Service) CreateConversation(conv *models.Conversation) (conversations []models.ConversationsInfo, err error) {
	service.Database.CreateConversation(conv)
	return
}

func (service *Service) CheckConversation(id1, id2 int, type_obj string) (conversations []models.ConversationsInfo, err error) {
	err = service.Database.VerifyConversation(id1, id2, type_obj)
	return
}
