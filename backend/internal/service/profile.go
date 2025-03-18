package service

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
)

func (service *Service) GetProfile(profile *models.User, currentUser int) (err error) {
	status, _ := service.Database.GetUserPrivacyByID(profile.ID)
	if profile.ID != currentUser && status == "private" && !service.Database.IsFollow(profile.ID, currentUser) {
		// not all info
		err = fmt.Errorf("profile is private, follow to see")
		return

	} else {
		// all info
		err = service.Database.GetFullProfile(profile)
		if err != nil {
			return
		}
	}
	return
}

func (service *Service) SetAction(profile *models.Profile, currentUser int) (err error) {
	if profile.ID == currentUser {
		profile.Action = "edit"
	} else {
		var status string
		status, err = service.Database.GetFollowerStatus(profile.ID, currentUser)
		if err == sql.ErrNoRows {
			profile.Action = "follow"
			err = nil
		} else if err != nil {
			return
		}
		if status == "pending" {
			profile.Action = "pending"
		} else if status == "accepted" {
			profile.Action = "unfollow"
		}
	}
	return
}

func (service *Service) GetFollowCounts(profile *models.Profile) (err error) {
	profile.Followers, err = service.Database.GetFollowersCount(&profile.UserInfo)
	if err != nil {
		return
	}
	profile.Followings, err = service.Database.GetFollowingsCount(&profile.UserInfo)
	return
}

func (service *Service) ModifyProfile(profile *models.User) (err error) {
	
	if (profile.Privacy != "public" && profile.Privacy != "private"){
		err = fmt.Errorf("bad request")
		return 
	}

	err = service.Database.UpdateUserPrivacy(profile)
	return

}
