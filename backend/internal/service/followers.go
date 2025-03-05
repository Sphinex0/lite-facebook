package service

import (
	"fmt"

	"social-network/internal/models"
)

func (S *Service) CreateFollow(follow *models.Follower) (err error) {
	if follow.UserID == follow.Follower {
		err = fmt.Errorf("can't follow yourself sadly")
		return
	}

	status, err := S.Database.GetUserPrivacyByID(follow)
	if err != nil {
		return
	}

	if status == "public" {
		follow.Status = "accepted"
	} else {
		follow.Status = "pending"
	}

	err = S.Database.SaveFollow(follow)
	return
}
