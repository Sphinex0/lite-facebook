package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveArticle(article *models.Article) (err error) {
	args := utils.GetExecFields(article, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO articles
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	article.ID = int(id)

	return
}
