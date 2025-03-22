package service

import (
	"database/sql"
	"fmt"
	"slices"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (service *Service) CreateInvite(Invites models.Invite) (err error) {
	boolean := service.Database.IsCreatore(Invites.Receiver, Invites.GroupID)
	if boolean {
		Invites.Status = "pending"
		id := 0
		id, err = service.Database.Saveinvite(&Invites)
		if err != nil {
			return fmt.Errorf("bad request")
		}
		if id == 0 {
			if Invites.Sender == Invites.Receiver {
				return fmt.Errorf("bad request")
			}
			err = service.Database.SaveInvite(&Invites)
		}
	}else {
		resA := service.Database.IsFollow(Invites.Sender, Invites.Receiver)
		resB := service.Database.IsFollow(Invites.Receiver, Invites.Sender)
		if resA && resB {
			err= service.Database.SaveInvite(&Invites)
			return
		} else {
			return fmt.Errorf("not follow")
		}
	}
	
	return
}



func (service *Service) InviderDecision(Invites *models.Invite) (err error) {
	if Invites.Status == "accepted" {
		err = service.Database.AcceptInviteRequest(Invites)
		fmt.Println(err)
		if err != nil {
			return
		}
		// for create member
		mb := Invites.Sender
		var idIv int
		idIv, err = service.Database.CheckMember(mb, Invites.GroupID)
		if err != nil && err != sql.ErrNoRows {
			return
		}
		if idIv == 0 {
			mb = Invites.Receiver
		}
		// get Conv by group id
		var conv models.Conversation
		conv, err = service.Database.GetConvByGroupID(Invites.GroupID)
		if err != nil {
			return
		}
		// create member
		member := models.Member{
			Member:         mb,
			ConversationId: conv.ID,
		}
		err = service.CreateMember(member)
	} else if Invites.Status == "rejected" {
		err = service.Database.DeleteInvites(Invites)
	} else {
		err = fmt.Errorf("bad  xcfdsf request")
	}
	return
}

func (S *Service) AllInvites(id int) ([]models.Invite, error) {
	rows, err := S.Database.GetallInvite(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Invites []models.Invite
	for rows.Next() {
		var Invite models.Invite
		if err := rows.Scan(utils.GetScanFields(&Invite)...); err != nil {
			fmt.Println(err)
			return nil, err
		}
		Invites = append(Invites, Invite)
	}

	return Invites, nil
}


func (S *Service) AllMembers(id int) ([]models.Invite, error) {
	rows, err := S.Database.Getallmembers(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Invites []models.Invite
	for rows.Next() {
		var Invite models.Invite
		if err := rows.Scan(utils.GetScanFields(&Invite)...); err != nil {
			fmt.Println("ttttttt",err)
			return nil, err
		}
		Invites = append(Invites, Invite)
	}

	return Invites, nil
}

func (S *Service) Members(Invite []models.Invite) ([]models.User, error) {
	var members []models.User
	for _, m := range Invite {
		fmt.Println("m.Receiver",m.Receiver)
		row1, err := S.Database.GetUsers(m.Receiver)
		if err != nil {
			fmt.Println("fffffffffffff",err)
			return nil, fmt.Errorf("error getting receiver user with ID %d: %w", m.Receiver, err)
		}
		if !slices.Contains(members, row1){
			members = append(members, row1)
		}
		

		row2, err := S.Database.GetUsers(m.Sender)
		if err != nil {
			fmt.Println("hhhhhhhhhhhhhhhhhhh",err)
			return nil, fmt.Errorf("error getting sender user with ID %d: %w", m.Sender, err)
		}
		if !slices.Contains(members, row2){
			members = append(members, row2)
		}
	}
	return members, nil
}

