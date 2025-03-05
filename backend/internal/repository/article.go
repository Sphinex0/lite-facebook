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

func (data *Database) GetArticlParent(id int) (err error) {
	var parent any
	err = data.Db.QueryRow(`
		SELECT parent FROM articles WHERE id = ?
	`, id).Scan(&parent)
	if parent != nil {
		err = fmt.Errorf("this comment dosn't have a parent")
	}
	return
}


func (data *Database) SaveReaction(article *models.Article) (err error) {
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