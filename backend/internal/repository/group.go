package repository

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveGroup(Group *models.Group) (err error) {
	args := utils.GetExecFields(Group, "ID")
	resGroup, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO groups
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := resGroup.LastInsertId()
	Group.ID = int(id)
	resInvite, err := data.Db.Exec(`
    INSERT INTO invites (group_id, sender, receiver, status)
    VALUES (?, ?, ?, 'accepted')`,
		Group.ID, Group.Creator, Group.Creator)
	if err != nil {
		return err // Handle error
	}

	// Optionally, get the ID of the invite if needed
	idgroup, err := resInvite.LastInsertId()
	if err != nil {
		return err // Handle error
	}
	var Invite models.Invite
	Invite.ID = int(idgroup)

	return
}

func (data *Database) Getallgroup() (*sql.Rows, error) {
	res, err := data.Db.Query(`SELECT * FROM groups`)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (data *Database) GetGroupById(id int) *sql.Row {
	res := data.Db.QueryRow(`SELECT * FROM groups Where id =?`, id)

	return res
}

func (data *Database) GetCreatorGroup(group_ID int, IdUser int) (bool, error) {
	res := data.Db.QueryRow(`SELECT creator FROM groups Where id =?`, group_ID)
	var id int

	err := res.Scan(&id)
	if err != nil {
		return false, err
	}
	if id == IdUser {
		return true, nil
	}
	return false, fmt.Errorf("not create group")
}

func (data *Database) Getmember(id int) (*sql.Rows, error) {
	query := `SELECT group_id FROM invites WHERE (sender = ? OR receiver = ?) AND status = "accepted"`

	return data.Db.Query(query, id, id)
}

func (data *Database) TypeUserInvate(id int, group_id int) (string, error) {
	types := ""
	query := `SELECT status FROM invites WHERE (sender = ? OR receiver = ?) AND group_id=? `

	err := data.Db.QueryRow(query, id, id, group_id).Scan(&types)
	if err != nil {
		if err == sql.ErrNoRows {
			return "join", err
		}
	}
	log.Println(types)
	return types, nil
}
