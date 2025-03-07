package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) GetInvites(id int) *sql.Row {
	res := data.Db.QueryRow(`SELECT creator FROM groups Where id =?`, id)

	return res
}
func (data *Database) ValidCreator(id int,IdSender int) bool {
	res := data.Db.QueryRow(`SELECT creator FROM groups Where id =?`, id)
	var creator int
	err := res.Scan(&creator)
	if err != nil{
		return false
	}

	if IdSender==creator {
		return true
	}
	return false
}

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
