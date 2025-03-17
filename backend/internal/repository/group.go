package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveGroup(Group *models.Group) (err error) {
	args := utils.GetExecFields(Group, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO groups
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	Group.ID = int(id)

	return
}

func (data *Database) Getallgroup() (*sql.Rows, error) {
	res, err := data.Db.Query(`SELECT * FROM groups`)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (data *Database) GetGroupById(id int) *sql.Row{
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



func (data *Database) Getmember(id int) (*sql.Rows,error){
	query := `SELECT group_id FROM invites WHERE (sender = ? OR receiver = ?) AND status = "accepted"`

    return data.Db.Query(query, id, id)

}