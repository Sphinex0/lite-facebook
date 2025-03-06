package service

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
)

func (S *Service) Follow(follow *models.Follower) (err error) {
	if follow.UserID == follow.Follower {
		err = fmt.Errorf("can't follow yourself sadly")
		return
	}

	followExist, err := S.Database.GetFollow(follow)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if followExist {
		err = S.Database.DeleteFollow(follow)
		if err != nil {
			return
		}
	} else {
		var status string
		status, err = S.Database.GetUserPrivacyByID(follow)
		if err != nil {
			return
		}
		if status == "public" {
			follow.Status = "accepted"
		} else {
			follow.Status = "pending"
		}

		err = S.Database.SaveFollow(follow)
	}
	return
}

func (S *Service) FollowDecision(follow *models.Follower) (err error) {
	
	if follow.Status == "accepted" {
		err = S.Database.AcceptFollowRequest(follow)
	} else if follow.Status == "rejected" {
		err = S.Database.DeleteFollow(follow)
	} else {
		err = fmt.Errorf("bad request")
	}

	return
}
