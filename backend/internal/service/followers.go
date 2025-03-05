package service

import "social-network/internal/models"

func (S *Service) CreateFollow(follow *models.Follower) (err error) {
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
