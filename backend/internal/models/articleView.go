package models

type ArticleView struct {
	UserInfo      UserInfo `json:"user_info"`
	Article       Article  `json:"article"`
	Likes         int      `json:"likes"`
	DisLikes      int      `json:"disLikes"`
	CommentsCount int      `json:"comments_count"`
	Like          int      `json:"like"`
}

type Data struct {
	Before  int `json:"before"`
	Parent  int `json:"parent"`
	GroupID int `json:"group_id"`
}
