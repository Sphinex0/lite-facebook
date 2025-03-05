package repository

import (
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

func (data *Database) GetUserPrivacyByID(follow *models.Follower) (status string, err error) {
	row := data.Db.QueryRow(`
        SELECT privacy 
		FROM users
		WHERE id = ?
    `,
	follow.UserID) 
	err = row.Scan(&status)	

	return
}


