package service

import "social-network/internal/models"

func (service *Service) CreateMember(member models.Member) (err error) {
	err = service.Database.SaveMember(member)
	return
}
// CheckMember
func (service *Service) CheckMember(mb, GroupID int) (id int, err error) {
	id, err = service.Database.CheckMember(mb, GroupID)
	return
}