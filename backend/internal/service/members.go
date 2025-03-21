package service

import "social-network/internal/models"

func (service *Service) CreateMember(member models.Member) (err error) {
	err = service.Database.SaveMember(member)
	return
}