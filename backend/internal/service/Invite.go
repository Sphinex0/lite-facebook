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

func (service *Service) InviderDecision(Invites *models.Invite) (err error) {
	
	if Invites.Status == "accepted" {
		err = service.Database.AcceptInviteRequest(Invites)
	} else if Invites.Status == "rejected" {
		err = service.Database.DeleteInvites(Invites)
	} else {
		err = fmt.Errorf("bad request")
	}

	return
}

