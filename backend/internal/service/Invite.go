package service

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (service *Service) CreateInvite(Invites models.Invite) (err error) {
	resA := service.Database.IsFollow(Invites.Sender, Invites.Receiver)
	resB := service.Database.IsFollow(Invites.Receiver, Invites.Sender)
	if resA && resB {

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
		} else {
			Invites.ID = id
			err = service.Database.DeleteInvites(&Invites)
		}
	} else {
		return fmt.Errorf("not follow")
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
		//member := models.Member{                                        
		//	Member : Group.Creator ,                                
		//	ConversationId : conv.ID ,
//}                                                       
//err = S.CreateMember()
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
			fmt.Println(err)
			return nil, err
		}
		Invites = append(Invites, Invite)
	}

	return Invites, nil
}


func (S *Service) Members(Invite []models.Invite) (map[int]models.User, error) {
	member:= make(map[int]models.User)
	for _, m := range Invite {
        row1, err := S.Database.GetUsers(m.Receiver)
        if err != nil {
            fmt.Println(err)
            return nil, fmt.Errorf("error getting receiver user with ID %d: %w", m.Receiver, err)
        }
        member[m.Receiver] = row1

        row2, err := S.Database.GetUsers(m.Sender)
        if err != nil {
            fmt.Println(err)
            return nil, fmt.Errorf("error getting sender user with ID %d: %w", m.Sender, err)
        }
        member[m.Sender] = row2
    }
	return member ,nil
}