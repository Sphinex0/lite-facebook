package repository

import (
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
