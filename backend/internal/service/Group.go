package service

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (S *Service) GreatedGroup(Group *models.Group) (err error) {
	err = S.Database.SaveGroup(Group)
	return
}

func (S *Service) AllGroups(Group *[]models.Group) ([]models.Group, error) {
	rows, err := S.Database.Getallgroup()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(utils.GetScanFields(&Group)); err != nil {
			fmt.Println(err)
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (S *Service) GetGroupsById(Group *models.Group, Id int) (*models.Group, error) {
	row := S.Database.GetGroupById(Id)
	if row == nil {
		return nil, fmt.Errorf("no group found with ID: %s", Id)
	}

	// Scan the row into the Group struct
	if err := row.Scan(utils.GetScanFields(&Group)); err != nil {
		return nil, fmt.Errorf("error scanning group data: %v", err)
	}

	return Group, nil
}
