package models

type ArticleView struct {
	UserInfo      UserInfo `json:"user_info"`
	Article       Article  `json:"article"`
	Likes         int      `json:"likes"`
	DisLikes      int      `json:"disLikes"`
	CommentsCount int      `json:"comments_count"`
	Like          int      `json:"like"`
	GroupName     *string  `json:"group_name"`
	GroupImage    *string  `json:"group_image"`
}

type Data struct {
	Before         int `json:"before"`
	Parent         int `json:"parent"`
	GroupID        int `json:"group_id"`
	ConversationID int `json:"conversation_id"`
	UserID int `json:"user_id"`
}
