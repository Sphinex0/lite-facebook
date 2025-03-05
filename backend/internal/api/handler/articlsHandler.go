package handler

import (
	"fmt"
	"net/http"
	"strings"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.User)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized)
		return
	}

	var article models.Article
	article.Content = strings.TrimSpace(r.FormValue("content"))

	fmt.Println(article)
}

func (Handler *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) AddReaction(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
