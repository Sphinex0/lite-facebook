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

func (data *Database) SaveReaction(like *models.Like) (err error) {
	args := utils.GetExecFields(like, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO likes
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	like.ID = int(id)

	return
}

func (data *Database) DeleteReaction(id int) (err error) {
	_, err = data.Db.Exec(`
		DELETE FROM likes 
		WHERE id = ?
	`, id)
	return
}

func (data *Database) UpdateReaction(id, like int) (err error) {
	_, err = data.Db.Exec(`
		UPDATE likes SET like = ?
		WHERE id = ?
	`, like, id)
	return
}

func (data *Database) GetReaction(user_id, article_id int) (id, like int, err error) {
	err = data.Db.QueryRow(`
		SELECT id , like 
		FROM likes 
		WHERE article_id = ? AND user_id = ?
	`, article_id, user_id).Scan(&id, &like)

	return
}
