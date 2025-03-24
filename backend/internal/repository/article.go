package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveArticle(article *models.Article, users []string) (err models.Error) {
	var res sql.Result
	args := utils.GetExecFields(article, "ID")
	res, err.Err = data.Db.Exec(fmt.Sprintf(`
		INSERT INTO articles
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err.Err != nil {
		return
	}
	var id int64
	id, err.Err = res.LastInsertId()
	article.ID = int(id)

	if article.Parent == nil {
		for _, user := range users {
			_, err.Err = data.Db.Exec(`
				INSERT INTO
					permited_users
				VALUES
					(NULL, ?, ?);
			`, id, user)
		}
	}

	return
}

func (data *Database) GetArticlParent(id int) (err models.Error) {
	var parent any
	err.Err = data.Db.QueryRow(`
		SELECT parent FROM articles WHERE id = ?
	`, id).Scan(&parent)
	if parent != nil {
		err.Err = fmt.Errorf("this comment dosn't have a parent")
	}
	return
}

func (data *Database) SaveReaction(like *models.Like) (err models.Error) {
	args := utils.GetExecFields(like, "ID")
	var res sql.Result
	res, err.Err = data.Db.Exec(fmt.Sprintf(`
		INSERT INTO likes
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err.Err != nil {
		return
	}
	var id int64
	id, err.Err = res.LastInsertId()
	if err.Err != nil {
		return
	}
	like.ID = int(id)

	return
}

func (data *Database) DeleteReaction(id int) (err models.Error) {
	_, err.Err = data.Db.Exec(`
		DELETE FROM likes 
		WHERE id = ?
	`, id)
	return
}

func (data *Database) UpdateReaction(id, like int) (err models.Error) {
	_, err.Err = data.Db.Exec(`
		UPDATE likes SET like = ?
		WHERE id = ?
	`, like, id)
	return
}

func (data *Database) GetReaction(user_id, article_id int) (id, like int, err models.Error) {
	err.Err = data.Db.QueryRow(`
		SELECT id , like 
		FROM likes 
		WHERE article_id = ? AND user_id = ?
	`, article_id, user_id).Scan(&id, &like)

	return
}

func (data *Database) GetPosts(id, before int) (article_views []models.ArticleView, err models.Error) {
	query := `
		SELECT DISTINCT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id), 0) as like
		FROM
			article_view A
			LEFT JOIN followers F ON A.user_id = F.user_id
			AND F.follower = ?
			AND F.status = 'accepted'
			LEFT JOIN permited_users P ON A.id = P.article_id
			AND P.user_id = ?
			LEFT JOIN invites I ON A.group_id = I.group_id AND I.status = "accepted" AND (I.sender = ? OR I.receiver = ?)
		WHERE
			parent IS NULL
			AND (
				A.privacy = 'public' AND (
					A.group_id IS NULL
					OR I.group_id IS NOT NULL
				)
				OR (
					A.user_id = ?
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
	var rows *sql.Rows
	rows, err.Err = data.Db.Query(query, id, id, id, id, id, id, before)
	if err.Err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab,
			&article_view.Likes,
			&article_view.DisLikes,
			&article_view.CommentsCount,
			&article_view.GroupName,
			&article_view.GroupImage,
			&article_view.Like,
		)

		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			return
		}
		article_views = append(article_views, article_view)
	}
	return
}

func (data *Database) GetPostsByUserId(id, user_id, before int) (article_views []models.ArticleView, err models.Error) {
	query := `
		SELECT DISTINCT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id), 0) as like
		FROM
			article_view A
			LEFT JOIN followers F ON A.user_id = F.user_id
			AND F.follower = ?
			AND F.status = 'accepted'
			LEFT JOIN permited_users P ON A.id = P.article_id
			AND P.user_id = ?
			LEFT JOIN invites I ON A.group_id = I.group_id AND I.status = "accepted" AND (I.sender = ? OR I.receiver = ?)
		WHERE
			parent IS NULL
			AND A.user_id = ?
			AND (
				A.privacy = 'public' AND (
					A.group_id IS NULL
					OR I.group_id IS NOT NULL
				)
				OR (
					A.user_id = ?
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

	var rows *sql.Rows
	rows, err.Err = data.Db.Query(query, id, id, id, id, id, user_id, id, before)
	if err.Err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab,
			&article_view.Likes,
			&article_view.DisLikes,
			&article_view.CommentsCount,
			&article_view.GroupName,
			&article_view.GroupImage,
			&article_view.Like,
		)

		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			return
		}
		article_views = append(article_views, article_view)
	}
	return
}

func (data *Database) GetComments(id, before, parent int) (article_views []models.ArticleView, err models.Error) {
	query := `
		SELECT DISTINCT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id),0) as like
		FROM
			article_view A
		WHERE
			A.parent = ? AND A.created_at < ?
        ORDER BY A.created_at DESC
        LIMIT 10
	`
	var rows *sql.Rows
	rows, err.Err = data.Db.Query(query, id, parent, before)
	if err.Err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab,
			&article_view.Likes,
			&article_view.DisLikes,
			&article_view.CommentsCount,
			&article_view.GroupName,
			&article_view.GroupImage,
			&article_view.Like,
		)

		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			return
		}
		article_views = append(article_views, article_view)
	}
	return
}

func (data *Database) VerifyGroupByID(group_id, id int) (err models.Error) {
	query := `
		SELECT id
		FROM invites
		WHERE group_id = ?
		AND status = "accepted"
		AND (
			sender = ? OR receiver = ?
		)
	`
	var result int
	err.Err = data.Db.QueryRow(query, group_id, id, id).Scan(&result)
	return
}

func (data *Database) GetPostsByGroup(id, group_id, before int) (article_views []models.ArticleView, err models.Error) {
	query := `
		SELECT DISTINCT
			A.*,
			COALESCE((SELECT like FROM likes L WHERE L.user_id = ? AND L.article_id = A.id),0) as like
		FROM
			article_view A
		WHERE
			A.group_id = ? AND A.parent IS NULL AND A.created_at < ?
        ORDER BY A.created_at DESC
        LIMIT 10
	`
	var rows *sql.Rows
	rows, err.Err = data.Db.Query(query, id, group_id, before)
	if err.Err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article_view models.ArticleView
		tab := utils.GetScanFields(&article_view.UserInfo)
		tab = append(tab, utils.GetScanFields(&article_view.Article)...)
		tab = append(tab,
			&article_view.Likes,
			&article_view.DisLikes,
			&article_view.CommentsCount,
			&article_view.GroupName,
			&article_view.GroupImage,
			&article_view.Like,
		)

		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			return
		}
		article_views = append(article_views, article_view)
	}
	return
}
