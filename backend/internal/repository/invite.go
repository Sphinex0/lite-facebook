package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveInvite(Invite *models.Invite) (err error) {
	args := utils.GetExecFields(Invite, "ID")

	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO Invites
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	Invite.ID = int(id)

	return
}

func (data *Database) AcceptInviteRequest(Invite *models.Invite) error {
	res, err := data.Db.Exec(`
        UPDATE Invites
		SET status = "accepted"
		WHERE id = ?
		AND status = "pending"
    `,
		Invite.ID)
	if err != nil {
		return err
	}
	res1, err1 := res.RowsAffected()
	if res1 == 0 || err1 != nil {
		return fmt.Errorf("not fount")
	}
	return err
}

func (data *Database) DeleteInvites(Invite *models.Invite) (err error) {
	_, err = data.Db.Exec(`
        DELETE 
		FROM Invites
		WHERE id = ?
    `,
		Invite.ID)

	return
}

func (data *Database) Saveinvite(Invite *models.Invite) (int, error) {
	res := data.Db.QueryRow(`SELECT id FROM invites WHERE group_id =? AND (sender = ? OR receiver =?) `, Invite.GroupID, Invite.Sender, Invite.Sender)
	var id int

	err := res.Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("bad request")
	}

	return id, nil
}
