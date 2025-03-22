package service

import (
	"database/sql"
	"fmt"
	"log"
	"slices"

	"social-network/internal/models"
)

func (S *Service) Follow(follow *models.Follower) (err error) {
	if follow.UserID == follow.Follower {
		err = fmt.Errorf("can't follow yourself sadly")
		return
	}


	var notification models.Notification

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
		status, err = S.Database.GetUserPrivacyByID(follow.UserID)
		if err != nil {
			log.Println("error checking targeted user privacy")

			return
		}
		if status == "public" {
			follow.Status = "accepted"
			notification.Type = "follow"
		} else {
			follow.Status = "pending"
			notification.Type = "follow-request"
		}

		err = S.Database.SaveFollow(follow)
		if err != nil {
			log.Println("error saving the follow")
		} else {
			if follow.Status == "accepted" {
				// chaeck if there is a conversation between the two
				conv, err1 := S.Database.GetConversationByUsers(follow.Follower, follow.UserID)
				if err1 != nil && err1 != sql.ErrNoRows {
					log.Println("error getting the conversation")
					err = err1
					return
				}
				if err1 == sql.ErrNoRows {
					conv = models.Conversation{
						Entitie_one:      follow.Follower,
						Entitie_two_user: &follow.UserID,
						Type:             "private",
					}
					S.Database.CreateConversation(&conv)

					ConvSubMu.Lock()
					defer ConvSubMu.Unlock()

					UserConnMu.RLock()
					defer UserConnMu.RUnlock()

					for _, userID := range []int{follow.Follower, follow.UserID} {
						if _, ok := UserConnections[userID]; ok {
							if !slices.Contains(ConvSubscribers[conv.ID], userID) {
								ConvSubscribers[conv.ID] = append(ConvSubscribers[conv.ID], userID)
							}
						}
					}
				}
			}
		}
		notification.InvokerID = follow.Follower
		notification.UserID = follow.UserID
		
		fmt.Println("rrrr")
		S.AddNotification(notification)

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
		if err == nil {
			conv := models.Conversation{
				Entitie_one:      follow.Follower,
				Entitie_two_user: &follow.UserID,
				Type:             "private",
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

func (S *Service) GetFollowRequests(user *models.UserInfo, before int, currentUser int) (requesters []models.UserInfo, err error) {
	status, _ := S.Database.GetUserPrivacyByID(user.ID)
	if user.ID != currentUser && status == "private" && !S.Database.IsFollow(user.ID, currentUser) {
		// not all info
		err = fmt.Errorf("profile is private, follow to see")
		return

	}
	requesters, err = S.Database.GetFollowRequests(user, before)
	return
}

func (S *Service) GetFollowers(user *models.UserInfo, before int, currentUser int) (followers []models.FollowWithUser, err error) {
	status, _ := S.Database.GetUserPrivacyByID(user.ID)
	if user.ID != currentUser && status == "private" && !S.Database.IsFollow(user.ID, currentUser) {
		// not all info
		err = fmt.Errorf("profile is private, follow to see")
		return

	}
	followers, err = S.Database.GetFollowers(user, before)
	return
}

func (S *Service) GetFollowings(user *models.UserInfo, before int, currentUser int) (followings []models.FollowWithUser, err error) {
	status, _ := S.Database.GetUserPrivacyByID(user.ID)
	if user.ID != currentUser && status == "private" && !S.Database.IsFollow(user.ID, currentUser) {
		// not all info
		err = fmt.Errorf("profile is private, follow to see")
		return

	}
	followings, err = S.Database.GetFollowings(user, before)
	return
}
