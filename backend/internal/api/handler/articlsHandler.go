package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) HandelCreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var article models.Article
	article.UserID = user.ID
	article.Content = strings.TrimSpace(r.FormValue("content"))
	article.Privacy = strings.TrimSpace(r.FormValue("privacy"))
	article.CreatedAt = int(time.Now().Unix())
	article.ModifiedAt = article.CreatedAt
	GroupID, _ := strconv.Atoi(r.FormValue("group_id"))
	if GroupID != 0 {
		/// select
		article.GroupID = &GroupID
	}
	parent, _ := strconv.Atoi(r.FormValue("parent"))
	if parent != 0 {
		/// select
		article.Parent = &parent
	}
	// fmt.Println(article)
	if err := Handler.Service.CreateArticle(&article); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
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
