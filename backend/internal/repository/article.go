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

func (data *Database) GetPosts(id, before int) (article_views []models.ArticleView, err error) {
	query := `
		SELECT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id),0) as like
		FROM
			article_view A
			LEFT JOIN followers F ON A.user_id = F.user_id
			AND F.follower = ?
			AND F.status = 'accepted'
			LEFT JOIN permited_users P ON A.id = P.article_id
			AND P.user_id = ?
		WHERE
			parent IS NULL
			AND (
				A.privacy = 'public'
				OR (
					A.id = ?
				)
				OR (
					A.privacy = 'almost_private'
					AND F.id IS NOT NULL
				)
				OR (
					A.privacy = 'private'
					AND P.user_id IS NOT NULL
				)
    		)
			AND A.created_at < ?

        ORDER BY A.created_at DESC
        LIMIT 10
	`

	rows, err := data.Db.Query(query, id, id, id, id, before)
	if err != nil {
		return
	}
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab, &article_view.Likes, &article_view.DisLikes, &article_view.CommentsCount, &article_view.Like)

		err1 := rows.Scan(tab...)
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		article_views = append(article_views, article_view)
	}
	rows.Close()
	return
}

func (data *Database) GetComments(id, before, parent int) (article_views []models.ArticleView, err error) {
	query := `
		SELECT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id),0) as like
		FROM
			article_view A
		WHERE
			A.parent = ? AND A.created_at < ?
        ORDER BY A.created_at DESC
        LIMIT 10
	`

	rows, err := data.Db.Query(query, id, parent, before)
	if err != nil {
		return
	}
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab, &article_view.Likes, &article_view.DisLikes, &article_view.CommentsCount, &article_view.Like)

		err1 := rows.Scan(tab...)
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		article_views = append(article_views, article_view)
	}
	rows.Close()
	return
}
