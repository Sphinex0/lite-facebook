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
