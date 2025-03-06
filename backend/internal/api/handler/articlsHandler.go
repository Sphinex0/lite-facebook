package handler

import (
	"fmt"
	"log"
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
		err := Handler.Service.VerifyParent(parent)
		if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		article.Parent = &parent
	}
	if err := Handler.Service.CreateArticle(&article); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (Handler *Handler) HandelGetArticles(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) HandelCreateReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var like models.Like
	err := utils.ParseBody(r, &like)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	like.UserID = user.ID

	err = Handler.Service.CreateReaction(&like)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (Handler *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
