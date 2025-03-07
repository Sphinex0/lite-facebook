package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveFollow(follow *models.Follower) (err error) {
	args := utils.GetExecFields(follow, "ID")

	_, err = data.Db.Exec(fmt.Sprintf(`
        INSERT INTO followers
        VALUES (NULL, %v) 
    `, utils.Placeholders(len(args))),
		args...)

	return
}

func (data *Database) GetFollow(follow *models.Follower) (followExist bool, err error) {
	row := data.Db.QueryRow(`
        SELECT *
		FROM followers
		WHERE user_id = ?
		AND follower = ?
    `,
		follow.UserID,
		follow.Follower)

	err = row.Scan(&follow)

	if err == nil {
		followExist = true
	} else if err == sql.ErrNoRows {
		followExist = false
	}

	return
}

func (data *Database) DeleteFollow(follow *models.Follower) (err error) {
	_, err = data.Db.Exec(`
        DELETE 
		FROM followers
		WHERE id = ?
    `,
		follow.ID)

	return
}

func (data *Database) GetFollowers(user *models.UserInfo) (followers []models.UserInfo, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image
		FROM followers f
		JOIN users u
		ON f.user_id = u.id
		WHERE user_id = ?
		AND status = "accepted"
    `,
		user.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		var user models.UserInfo
		if err = rows.Scan(&user); err != nil {
			return
		}
		followers = append(followers, user)
	}

	return
}

func (data *Database) GetFollowByUser(user int, creator int) (err error) {
	var id int
	err = data.Db.QueryRow(`
        SELECT id
		FROM followers 
		WHERE user_id = ?
		AND follower = ?
		AND status = "accepted"
    `,
		creator, user).Scan(&id)

	return
}

func (data *Database) GetFollowings(user *models.UserInfo) (followers []models.UserInfo, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image
		FROM followers f
		JOIN users u
		ON f.follower = u.id
		WHERE f.follower = ?
		AND status = "accepted"
    `,
		user.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		var user models.UserInfo
		if err = rows.Scan(&user); err != nil {
			return
		}
		followers = append(followers, user)
	}

	return
}

func (data *Database) GetFollowRequests(user *models.UserInfo) (Requests []models.Follower, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT *
		FROM followers
		WHERE user_id = ?
		AND status = "pending"
    `,
		user.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		var follow models.Follower
		if err = rows.Scan(&follow); err != nil {
			return
		}
		Requests = append(Requests, follow)
	}

	return
}

func (data *Database) AcceptFollowRequest(follow *models.Follower) (err error) {
	_, err = data.Db.Exec(`
        UPDATE followers
		SET status = "accepted"
		WHERE id = ?
		AND status = "pending"
    `,
		follow.ID)

	return
}

func (data *Database) GetUserPrivacyByID(follow *models.Follower) (status string, err error) {
	err = data.Db.QueryRow(`
        SELECT privacy 
		FROM users
		WHERE id = ?
    `,
		follow.UserID).Scan(&status)

	return
}
