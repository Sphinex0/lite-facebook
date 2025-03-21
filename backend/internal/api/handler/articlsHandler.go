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
)

func (Handler *Handler) HandelCreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var article models.Article
	article.UserID = user.ID
	article.Content = strings.TrimSpace(r.FormValue("content"))
	article.Privacy = strings.TrimSpace(r.FormValue("privacy"))
	article.CreatedAt = int(time.Now().UnixMilli())
	article.ModifiedAt = article.CreatedAt
	GroupID, _ := strconv.Atoi(r.FormValue("group_id"))
	var users []string
	if article.Privacy == "private" {
		users = r.Form["users"]
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		article.Image, err = utils.StoreThePic("../front-end/public/posts", file, handler)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
	}

	parent, _ := strconv.Atoi(r.FormValue("parent"))
	if parent != 0 {
		/// select
		err := Handler.Service.VerifyParent(parent)
		if err.Err != nil {
			utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		article.Parent = &parent
		article.Privacy = "public"
	} else if GroupID != 0 {
		/// select
		// err := Handler.Service.VerifyGroup(GroupID, user.ID)
		// if err.Err != nil {
		// 	utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		// 	return
		// }
		_, err := Handler.Service.CheckMember(GroupID, user.ID)
		if err != nil {
			utils.WriteJson(w, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}
		article.GroupID = &GroupID
		article.Privacy = "public"

	}
	if err := Handler.Service.CreateArticle(&article, users, user.ID); err.Err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	utils.WriteJson(w, http.StatusCreated, article)
}

func (Handler *Handler) HandelGetPosts(w http.ResponseWriter, r *http.Request) {
	user, data, err := Handler.AfterGet(w, r)
	if err.Err != nil {
		return
	}
	article_views, err := Handler.Service.FetchPosts(user.ID, data.Before)
	if err.Err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.WriteJson(w, http.StatusOK, article_views)
}

func (Handler *Handler) HandelGetPostsByGroup(w http.ResponseWriter, r *http.Request) {
	user, data, err := Handler.AfterGet(w, r)
	if err.Err != nil {
		return
	}
	err = Handler.Service.VerifyGroup(data.GroupID, user.ID)
	if err.Err != nil {
		utils.WriteJson(w, http.StatusForbidden, http.StatusText(http.StatusForbidden))
		return
	}
	article_views, err := Handler.Service.FetchPostsByGroup(user.ID, data.GroupID, data.Before)
	if err.Err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.WriteJson(w, http.StatusOK, article_views)
}

func (Handler *Handler) HandelGetComments(w http.ResponseWriter, r *http.Request) {
	user, data, err := Handler.AfterGet(w, r)
	if err.Err != nil {
		return
	}
	article_views, err := Handler.Service.FetchComments(user.ID, data.Before, data.Parent)
	if err.Err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.WriteJson(w, http.StatusOK, article_views)
}

func (Handler *Handler) HandelCreateReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var like models.Like
	var err models.Error
	err.Err = utils.ParseBody(r, &like)
	if err.Err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	like.UserID = user.ID

	err = Handler.Service.CreateReaction(&like)
	if err.Err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.WriteJson(w, http.StatusCreated, like)
}

func (Handler *Handler) AfterGet(w http.ResponseWriter, r *http.Request) (user models.UserInfo, data models.Data, err models.Error) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		err.Err = fmt.Errorf("err in methode")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		err.Err = fmt.Errorf("Unauthorized")
		return
	}

	err.Err = utils.ParseBody(r, &data)
	if err.Err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if data.Before == 0 {
		data.Before = int(time.Now().UnixMilli())
	}
	return
}
