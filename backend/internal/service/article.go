package service

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strconv"

	"social-network/internal/models"
)

func (service *Service) CreateArticle(article *models.Article, users []string, id int) (err models.Error) {
	if article.Content == "" {
		err.Err = fmt.Errorf("err in content")
		return
	}
	privacies := []string{"public", "private", "almost_private"}
	isAllowedPrivacy := slices.Contains(privacies, article.Privacy)
	if !isAllowedPrivacy {
		article.Privacy = "public"
	}
	var followerIds []int

	followerIds, err = service.Database.GetFollowersIds(id)
	if err.Err != nil {
		return
	}

	for _, user := range users {
		id, err.Err = strconv.Atoi(user)
		if err.Err != nil {
			return
		}
		if !slices.Contains(followerIds, id) {
			err.Err = fmt.Errorf("error in user")
			return
		}
	}

	err = service.Database.SaveArticle(article, users)
	return
}

func (service *Service) VerifyParent(parent int) (err models.Error) {
	err = service.Database.GetArticlParent(parent)
	return
}

func (service *Service) VerifyGroup(group_id, id int) (err models.Error) {
	err = service.Database.VerifyGroupByID(group_id, id)
	return
}

func (service *Service) CreateReaction(like *models.Like) (err models.Error) {
	if (like.Like != 1 && like.Like != -1) || like.ArticleID == 0 {
		err.Err = errors.New("invalid Like value")
		return
	}

	id, existsLike, err := service.Database.GetReaction(like.UserID, like.ArticleID)
	if err.Err == sql.ErrNoRows {
		err = service.Database.SaveReaction(like)
	} else if existsLike == like.Like {
		err = service.Database.DeleteReaction(id)
	} else if err.Err == nil {
		err = service.Database.UpdateReaction(id, like.Like)
	}
	return
}

func (service *Service) FetchPosts(id, before int) (article_views []models.ArticleView, err models.Error) {
	article_views, err = service.Database.GetPosts(id, before)
	return
}

func (service *Service) FetchPostsByGroup(id, group_id, before int) (article_views []models.ArticleView, err models.Error) {
	article_views, err = service.Database.GetPostsByGroup(id, group_id, before)
	return
}

func (service *Service) FetchComments(id, before, parent int) (article_views []models.ArticleView, err models.Error) {
	if parent == 0 {
		err.Err = fmt.Errorf("err in parent")
		return
	}
	article_views, err = service.Database.GetComments(id, before, parent)
	return
}
