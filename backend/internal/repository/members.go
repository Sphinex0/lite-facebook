package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveMember(member models.Member) (err error) {
	args := utils.GetExecFields(member, "ID")
	_, err = data.Db.Exec(fmt.Sprintf(`
        INSERT INTO members
        VALUES (NULL, %v)
    `, utils.Placeholders(len(args))), args...)
	return
}

func (data *Database) CheckMember(mb, GroupID int) (id int, err error) {
	err = data.Db.QueryRow(`
         SELECT id
         FROM members 
         WHERE member = ? AND group_id = ?
    `, mb, GroupID).Scan(&id)
	return
}
