package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveInvite(Invite *models.Invite) (err error) {
	fmt.Println(Invite)
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

func (data *Database) AcceptInviteRequest(Invite *models.Invite) (err error) {
	_, err = data.Db.Exec(`
        UPDATE Invites
		SET status = "accepted"
		WHERE id = ?
		AND status = "pending"
    `,
		Invite.ID)

	return
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
