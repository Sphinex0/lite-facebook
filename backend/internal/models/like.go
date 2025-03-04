package models

// Like represents a like on an article.
type Like struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ArticleID string `json:"article_id"`
	Like      string `json:"like"`
}
