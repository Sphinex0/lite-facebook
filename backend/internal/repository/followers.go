package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)



func (data *Database) GetFollowersCount(user *models.UserInfo) (count int, err error) {
	row := data.Db.QueryRow(`
        SELECT COUNT(*)
		FROM followers
		WHERE user_id = ?
		AND status = "accepted"
    `,
		user.ID)

	err = row.Scan(&count)

	return
}

func (data *Database) GetFollowingsCount(user *models.UserInfo) (count int, err error) {
	row := data.Db.QueryRow(`
        SELECT COUNT(*)
		FROM followers
		WHERE follower = ?
		AND status = "accepted"
    `,
		user.ID)

	err = row.Scan(&count)

	return
}

func (data *Database) SaveFollow(follow *models.Follower) (err error) {
	args := utils.GetExecFields(follow, "ID")

	_, err = data.Db.Exec(fmt.Sprintf(`
        INSERT INTO followers
        VALUES (NULL, %v) 
    `, utils.Placeholders(len(args))),
		args...)

	return
}

func (data *Database) IsFollow(user1 int, user2 int) (following bool) {
	var count int
	err := data.Db.QueryRow(`
        SELECT COUNT(*)
		FROM followers
		WHERE user_id = ?
		AND follower = ?
		AND status = "accepted"
    `,
		user1,
		user2).Scan(&count)
	if err == nil && count == 1 {
		following = true
	} else {
		following = false
	}

	return
}

func (data *Database) GetFollowerStatus(targetUser int, FollowerUser int) (status string, err error) {
	err = data.Db.QueryRow(`
        SELECT status
		FROM followers
		WHERE user_id = ?
		AND follower = ?
    `,
	targetUser,
	FollowerUser).Scan(&status)

	return
}

func (data *Database) GetFollow(follow *models.Follower) (err error) {
	row := data.Db.QueryRow(`
        SELECT *
		FROM followers
		WHERE user_id = ?
		AND follower = ?
    `,
		follow.UserID,
		follow.Follower)

	err = row.Scan(utils.GetScanFields(follow)...)

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

func (data *Database) GetFollowersIds(userID int) (followerIds []int, err models.Error) {
	var rows *sql.Rows
	rows, err.Err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image
		FROM followers f
		JOIN users u
		ON f.follower = u.id
		WHERE user_id = ?
		AND status = "accepted"

    `,
		userID)
	if err.Err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserInfo
		if err.Err = rows.Scan(utils.GetScanFields(&user)...); err.Err != nil {
			return
		}
		followerIds = append(followerIds, user.ID)
	}
	err.Err = rows.Err()

	return
}

func (data *Database) GetFollowers(user *models.UserInfo, before int) (followers []models.FollowWithUser, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image, f.created_at, f.modified_at
		FROM followers f
		JOIN users u
		ON f.follower = u.id
		WHERE user_id = ?
		AND status = "accepted"
		AND modified_at < ?
		ORDER BY modified_at DESC
		LIMIT 10
    `,
		user.ID,
		before)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var FollowUser models.FollowWithUser
		scanFields := utils.GetScanFields(&FollowUser.UserInfo)
		scanFields = append(scanFields, &FollowUser.CreatedAt, &FollowUser.ModifiedAt)
		if err = rows.Scan(scanFields...); err != nil {
			return
		}
		followers = append(followers, FollowUser)
	}
	err = rows.Err()

	return
}

func (data *Database) GetFollowings(user *models.UserInfo, before int) (followings []models.FollowWithUser, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image, f.created_at, f.modified_at
		FROM followers f
		JOIN users u
		ON f.user_id = u.id
		WHERE f.follower = ?
		AND status = "accepted"
		AND modified_at < ?
		ORDER BY modified_at DESC
		LIMIT 10
    `,
		user.ID,
		before)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var FollowUser models.FollowWithUser
		scanFields := utils.GetScanFields(&FollowUser.UserInfo)
		scanFields = append(scanFields, &FollowUser.CreatedAt, &FollowUser.ModifiedAt)
		if err = rows.Scan(scanFields...); err != nil {
			return
		}
		followings = append(followings, FollowUser)
	}
	err = rows.Err()

	return
}

// to get the follow without the status (used to modify old status)
func (data *Database) GetPendingFollowByUsers(follow *models.Follower) (err error) {
	err = data.Db.QueryRow(`
        SELECT id, user_id, follower
		FROM followers 
		WHERE user_id = ?
		AND follower = ?
		AND status = "pending"
		LIMIT 20
    `,
		follow.UserID, follow.Follower).Scan(&follow.ID, &follow.UserID, &follow.Follower)

	return
}

func (data *Database) GetFollowRequests(user *models.UserInfo, before int) (requesters []models.UserInfo, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
        SELECT u.id, u.nickname, u.first_name, u.last_name, u.image
		FROM followers f
		JOIN users u
		ON f.follower = u.id
		WHERE user_id = ?
		AND status = "pending"
		AND modified_at < ?
		LIMIT 20
    `,
		user.ID,
		before)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserInfo
		if err = rows.Scan(utils.GetScanFields(&user)...); err != nil {
			return
		}
		requesters = append(requesters, user)
	}
	err = rows.Err()

	return
}

func (data *Database) AcceptFollowRequest(follow *models.Follower) (err error) {
	_, err = data.Db.Exec(`
        UPDATE followers
		SET status = "accepted",
			modified_at = ?
		WHERE id = ?
    `,
		follow.ModifiedAt,
		follow.ID)

	return
}

func (data *Database) GetUserPrivacyByID(userID int ) (status string, err error) {
	err = data.Db.QueryRow(`
        SELECT privacy 
		FROM users
		WHERE id = ?
    `,
	userID).Scan(&status)

	return
}
