package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveEvent(Event *models.Event) (err error) {
	args := utils.GetExecFields(Event, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO Events
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	Event.ID = int(id)

	return
}
