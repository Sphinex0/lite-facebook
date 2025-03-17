package service

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/internal/models"
)

func (S *Service) Follow(follow *models.Follower) (err error) {
	if follow.UserID == follow.Follower {
		err = fmt.Errorf("can't follow yourself sadly")
		return
	}

	err = S.Database.GetFollow(follow)

	if err == nil {
		err = S.Database.DeleteFollow(follow)
		if err != nil {
			log.Println("error deleting the follow")

			return
		}
	} else if err == sql.ErrNoRows {
		err = nil
		var status string
		status, err = S.Database.GetUserPrivacyByID(follow)
		if err != nil {
			log.Println("error checking targeted user privacy")

			return
		}
		if status == "public" {
			follow.Status = "accepted"
		} else {
			follow.Status = "pending"
		}

		err = S.Database.SaveFollow(follow)
		if err != nil {
			log.Println("error saving the follow")
		}else {
			if follow.Status == "accepted" {
				conv := models.Conversation{
					Entitie_one: follow.Follower,
					Entitie_two_user: &follow.UserID,
					Type: "private",
				}
			
				S.Database.CreateConversation(&conv)
			}
		}

	}
	return
}

func (S *Service) FollowDecision(follow *models.Follower) (err error) {
	err = S.Database.GetPendingFollowByUsers(follow)
	if err != nil {
		log.Println("error finding the pending follow request")
		return
	}

	if follow.Status == "accepted" {
		err = S.Database.AcceptFollowRequest(follow)
		if err == nil{
			conv := models.Conversation{
				Entitie_one: follow.Follower,
				Entitie_two_user: &follow.UserID,
				Type: "private",
			}
		
			S.Database.CreateConversation(&conv)
		}
	} else if follow.Status == "rejected" {
		err = S.Database.DeleteFollow(follow)
	} else {
		log.Println("unkown status")
		err = fmt.Errorf("bad request")
	}

	return
}

func (S *Service) GetFollowRequests(user *models.UserInfo, before int) (requesters []models.UserInfo,err error) {
	requesters , err = S.Database.GetFollowRequests(user, before)
	return
}

func (S *Service) GetFollowers(user *models.UserInfo, before int) (followers []models.FollowWithUser,err error) {
	followers , err = S.Database.GetFollowers(user, before)
	return
}
func (S *Service) GetFollowings(user *models.UserInfo, before int) (followings []models.UserInfo,err error) {
	followings , err = S.Database.GetFollowings(user, before)
	return
}
