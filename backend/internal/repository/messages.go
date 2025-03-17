package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveMessage(msg *models.Message) (err error) {
	args := utils.GetExecFields(msg, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO messages
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	msg.ID = int(id)
	return
}
