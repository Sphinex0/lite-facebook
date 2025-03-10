package service

import (
	"social-network/internal/models"
)

func (service *Service) CreateEvent(Events models.Event) (err error) {
	valid ,err :=service.Database.GetCreatorGroup(Events.GroupID,Events.UserID)
	if err !=nil{
		return 
	} 
	if valid {
		err = service.Database.SaveEvent(&Events)
	}
	return
}
