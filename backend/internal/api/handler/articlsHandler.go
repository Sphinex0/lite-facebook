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

func (Handler *Handler) AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.User)
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
	GroupID , err := strconv.Atoi(r.FormValue("group_id"))
	if err != nil || GroupID == 0 {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	article.GroupID = GroupID
	parent , err := strconv.Atoi(r.FormValue("parent"))
	if err != nil || parent == 0 {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	article.Parent = parent
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
