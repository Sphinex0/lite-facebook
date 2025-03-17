package service

import (
	"time"

	"social-network/internal/models"
)

func (service *Service) CreateMessage(msg *models.Message) (err error) {
	msg.CreatedAt = int(time.Now().Unix())
	err = service.Database.SaveMessage(msg)
	return
}
