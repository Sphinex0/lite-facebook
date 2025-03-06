package models

type ArticleView struct {
	UserInfo      UserInfo `json:"user_info"`
	Article       Article  `json:"article"`
	Likes         int      `json:"likes"`
	DisLikes      int      `json:"disLikes"`
	CommentsCount int      `json:"comments_count"`
	Like          int      `json:"like"`
}
