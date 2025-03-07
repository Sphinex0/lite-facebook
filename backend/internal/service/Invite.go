package service

import (
	"fmt"

	"social-network/internal/models"
)

func (service *Service) CreateInvite(Invites models.Invite) (err error) {
	Invites.Status = "pending"
	if Invites.Sender == Invites.Receiver {
		return fmt.Errorf("bad request")
	}
	err = service.Database.SaveInvite(&Invites)

	return
}

