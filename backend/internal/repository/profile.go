package repository

import (
	"database/sql"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (database *Database) GetFullProfile(profile *models.User) (err error) {
	err = database.Db.QueryRow(`SELECT * FROM users  WHERE id = ? 
	`, profile.ID).Scan(utils.GetScanFields(profile)...)
	profile.Password = ""
	return
}

func (database *Database) UpdateUserPrivacy(profile *models.User) (err error) {
	_, err = database.Db.Exec(`UPDATE users SET  privacy = ? WHERE id = ? 
	`, profile.Privacy, profile.ID)
	return
}

func (data *Database) GetAllUsers(before int, currentUser int) (users []models.UserInfo, err error) {
	var rows *sql.Rows
	if before != 0 {
		rows, err = data.Db.Query(`
        SELECT id, nickname, first_name, last_name, image
		FROM users 
		WHERE id != ?
		AND id < ?
		ORDER BY id DESC
		LIMIT 10
    `,
			currentUser,
			before)
	} else {
		rows, err = data.Db.Query(`
        SELECT id, nickname, first_name, last_name, image
		FROM users u
		WHERE id != ?
		ORDER BY id DESC
		LIMIT 10
    `,
	currentUser)
	}

	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserInfo
		if err = rows.Scan(utils.GetScanFields(&user)...); err != nil {
			return
		}
		users = append(users, user)
	}
	err = rows.Err()

	return
}
