package service

import (
	"fmt"
	"log"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (S *Service) GreatedGroup(Group *models.Group) (err error) {
	err = S.Database.SaveGroup(Group)
	if err != nil {
		return
	}
	// create conv and member
	tm := int(time.Now().UnixMilli())
	conv := &models.Conversation{
		Entitie_one:       Group.Creator,
		Entitie_two_group: &Group.ID,
		Type:              "group",
		CreatedAt:         tm,
		ModifiedAt:        tm,
	}
	_, err = S.CreateConversation(conv)
	if err != nil {
		return
	}
	member := models.Member{
		Member:         Group.Creator,
		ConversationId: conv.ID,
	}
	err = S.CreateMember(member)
	return
}

func (S *Service) AllGroups() ([]models.Group, error) {
	rows, err := S.Database.Getallgroup()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(utils.GetScanFields(&group)...); err != nil {
			fmt.Println(err)
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (S *Service) GetGroupsById(Group *models.Group) (*models.Group, error) {
	row := S.Database.GetGroupById(Group.ID)
	if row == nil {
		return nil, fmt.Errorf("no group found with ID: %s", Group.ID)
	}

	// Scan the row into the Group struct
	if err := row.Scan(utils.GetScanFields(Group)...); err != nil {
		return nil, fmt.Errorf("error scanning group data: %v", err)
	}

	return Group, nil
}

func (S *Service) TypeInvate(id int, group_id int) (string, error) {
	types := ""
	types, err := S.Database.TypeUserInvate(id, group_id)
	if err != nil {
		if types == "" {
			return "", err
		}
	}
	return types, nil
}

func (S *Service) GetMemberById(GroupId int) ([]models.Group, error) {
	rows, err := S.Database.Getmember(GroupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)

	Group := []int{}
	for rows.Next() {
		var groupIDScan int
		if err := rows.Scan(&groupIDScan); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		Group = append(Group, groupIDScan)
	}

	log.Println(Group)
	var groups []models.Group

	for _, v := range Group {
		var group models.Group
		rowGroupe := S.Database.GetGroupById(v)
		if err := rowGroupe.Scan(utils.GetScanFields(&group)...); err != nil {
			return nil, fmt.Errorf("error scanning group data: %v", err)
		}
		groups = append(groups, group)
	}

	return groups, nil
}
