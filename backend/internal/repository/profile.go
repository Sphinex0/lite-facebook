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