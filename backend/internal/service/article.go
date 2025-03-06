package service

import (
	"database/sql"
	"errors"
	"slices"

	"social-network/internal/models"
)

func (service *Service) CreateArticle(article *models.Article) (err error) {
	privacies := []string{"public", "private", "almost_private"}
	isAllowedPrivacy := slices.Contains(privacies, article.Privacy)
	if !isAllowedPrivacy {
		article.Privacy = "public"
	}
	err = service.Database.SaveArticle(article)
	return
}

func (service *Service) VerifyParent(parent int) (err error) {
	err = service.Database.GetArticlParent(parent)
	return
}

func (service *Service) CreateReaction(like *models.Like) (err error) {
	if (like.Like != 1 && like.Like != -1) || like.ArticleID == 0 {
		return errors.New("invalid Like value")
	}

	id, existsLike, err := service.Database.GetReaction(like.UserID, like.ArticleID)
	if err == sql.ErrNoRows {
		err = service.Database.SaveReaction(like)
	} else if existsLike == like.Like {
		err = service.Database.DeleteReaction(id)
	} else if err == nil {
		err = service.Database.UpdateReaction(id, like.Like)
	}
	return
}
