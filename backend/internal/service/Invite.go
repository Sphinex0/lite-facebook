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

func (service *Service) InviderDecision(follow *models.Follower) (err error) {
	
	if follow.Status == "accepted" {
		err = service.Database.AcceptFollowRequest(follow)
	} else if follow.Status == "rejected" {
		err = service.Database.DeleteFollow(follow)
	} else {
		err = fmt.Errorf("bad request")
	}

	return
}

