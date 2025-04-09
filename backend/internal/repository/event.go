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

func (data *Database) GetallEvents(idGroup int) (*sql.Rows, error) {
	res, err := data.Db.Query(`SELECT * FROM Events WHERE group_id=? ORDER BY id DESC `, idGroup)
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
	res, err := data.Db.Query(`SELECT * FROM event_options WHERE event_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (data *Database) CheckEvent(EventId, UserId int) (bool, error) {
	var going bool

	err := data.Db.QueryRow(`SELECT going FROM event_options WHERE event_id = ? AND user_id = ?`, EventId, UserId).Scan(&going)

	return going, err
}

// GetEventCountsAndUserChoice fetches counts and user's choice for an event
func (data *Database) GetEventCountsAndUserChoice(eventID int, userID int) (int, int, sql.NullBool, error) {
	var goingCount, notGoingCount int
	var userGoing sql.NullBool

	query := `
        SELECT 
            COUNT(CASE WHEN going = true THEN 1 END) AS going_count,
            COUNT(CASE WHEN going = false THEN 1 END) AS not_going_count,
            (SELECT going FROM event_options WHERE event_id = ? AND user_id = ?) AS user_going
        FROM event_options
        WHERE event_id = ?
    `
	err := data.Db.QueryRow(query, eventID, userID, eventID).Scan(&goingCount, &notGoingCount, &userGoing)
	if err != nil {
		return 0, 0, sql.NullBool{}, err
	}
	return goingCount, notGoingCount, userGoing, nil
}

func (data *Database) UpdateOptionEvent(EventOption models.EventOption) error {
	_, err := data.Db.Exec(`UPDATE event_options SET going = ? WHERE event_id = ? AND user_id = ?`, EventOption.Going, EventOption.EventID, EventOption.UserID)
	return err
}
