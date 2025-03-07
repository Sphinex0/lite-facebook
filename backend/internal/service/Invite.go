package service

import (
	"fmt"

	"social-network/internal/models"
)

func (service *Service) CreateInviteByReceiver(IdGroup int, IdReciver int) (err error) {
	row := service.Database.GetInvites(IdGroup)
	if row == nil {
		return fmt.Errorf("no group found with ID: %s", IdGroup)
	}
	var IdSender int
	err = row.Scan(&IdSender)
	if err != nil {
		return fmt.Errorf("found with ID: %d", IdSender)
	}
	if IdSender == IdReciver {
		return fmt.Errorf("bad request")
	}
	var Invite models.Invite

	Invite.GroupID = IdGroup
	Invite.Sender = IdSender
	Invite.Receiver = IdReciver
	Invite.Status = "pending"
	fmt.Println(Invite)

	err = service.Database.SaveInvite(&Invite)

	return
}

func (service *Service) CreateInviteBySernder(IdGroup int, IdSender int, IdReciver int) (err error) {
	row := service.Database.ValidCreator(IdGroup, IdSender)
	if !row {
		return fmt.Errorf("no group found with ID: %s", IdGroup)
	}

	if IdSender == IdReciver {
		return fmt.Errorf("bad request")
	}

	var Invite models.Invite

	Invite.GroupID = IdGroup
	Invite.Sender = IdSender
	Invite.Receiver = IdReciver
	Invite.Status = "pending"
	err = service.Database.SaveInvite(&Invite)
	return
}
