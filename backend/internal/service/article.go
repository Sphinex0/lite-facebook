package service

import (
	"slices"
	"strings"

	"social-network/internal/models"
)

func (service *Service) CreateArticle(article *models.Article) (err error) {
	privacies := []string{"public", "private", "almost_private"}
	index := slices.IndexFunc(privacies, func(ext string) bool {
		return strings.Contains(article.Privacy, ext)
	})
	if index == -1 {
		article.Privacy = "public"
	}
	err = service.Database.SaveArticle(article)
	return
}

func (service *Service) VerifyParent(parent int) (err error) {
	err = service.Database.GetArticlParent(parent)
	return
}


func (service *Service) CreateReaction(article *models.Article) (err error) {
	privacies := []string{"public", "private", "almost_private"}
	index := slices.IndexFunc(privacies, func(ext string) bool {
		return strings.Contains(article.Privacy, ext)
	})
	if index == -1 {
		article.Privacy = "public"
	}
	err = service.Database.SaveReaction(article)
	return
}
