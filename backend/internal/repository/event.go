package repository

import (
	"database/sql"
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

func (data *Database) GetallEvents() (*sql.Rows, error) {
	res, err := data.Db.Query(`SELECT * FROM Events`)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (data *Database) GetEventById(id int) *sql.Row {
	res := data.Db.QueryRow(`SELECT * FROM Events Where id =?`, id)
	return res
}

func (data *Database) SaveOptionEvent(Event *models.EventOption) (err error) {
	args := utils.GetExecFields(Event, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO event_options
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

func (data *Database) OptionEvent(id int) (*sql.Rows, error) {
	res, err := data.Db.Query(`SELECT * FROM event_options WHERE even_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
