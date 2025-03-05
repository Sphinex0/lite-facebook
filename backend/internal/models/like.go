package models

// Like represents a like on an article.
type Like struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ArticleID int `json:"article_id"`
	Like      int `json:"like"`
}
