package service

func (service *Service) CreateMember(member models.member) (err error) {
	err = service.Database.SaveMember(member)
}