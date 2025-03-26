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
	fmt.Println(Invite)
	res, err := data.Db.Exec(`
        UPDATE Invites
		SET status = "accepted"
		WHERE sender = ?
		AND receiver = ?
		AND group_id = ?
		AND status = "pending"
    `,
		Invite.Sender,
		Invite.Receiver,
		Invite.GroupID)
	if err != nil {
		return err
	}
	res1, err1 := res.RowsAffected()
	fmt.Println(err1)
	fmt.Println(res1)
	if res1 == 0 || err1 != nil {
		return fmt.Errorf("not fount")
	}
	return err
}

func (data *Database) DeleteInvites(Invite *models.Invite) (err error) {
	_, err = data.Db.Exec(`
        DELETE 
		FROM Invites
		WHERE sender = ?
		AND receiver = ?
		AND group_id = ?
    `,
	Invite.Sender,
	Invite.Receiver,
	Invite.GroupID)

	return
}

func (data *Database) Saveinvite(Invite *models.Invite) (int, error) {
	res := data.Db.QueryRow(`
    SELECT id 
    FROM invites 
    WHERE group_id = ? AND sender = ? AND receiver = ?
`, Invite.GroupID, Invite.Sender, Invite.Receiver)
	var id int

	err := res.Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("bad request")
	}

	return id, nil
}

func (data *Database) IsCreatore(resever int , group_Id int) bool{
	var Create int
	err:= data.Db.QueryRow(`
	SELECT creator FROM groups WHERE id = ? `, group_Id).Scan(&Create)
	if err != nil {
		return false
	}
	if resever == Create {
		return true
	}
	return false
}

func (data *Database) GetallInvite(id int) (*sql.Rows, error) {
	res, err := data.Db.Query(`
	SELECT * FROM invites WHERE 
	(sender = ? OR receiver = ?) AND status = "pending"`, id, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (data *Database) Getallmembers(id int) (*sql.Rows, error) {
	res, err := data.Db.Query(`
	SELECT * FROM invites WHERE group_id = ? AND status = "accepted"`, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (data *Database) GetUsers(Id int) (models.User, error) {
	
	var user models.User
	row := data.Db.QueryRow("SELECT * FROM users WHERE id = ?", Id)
	if  err := row.Scan(utils.GetScanFields(&user)...); err != nil {
		fmt.Println("/api/invites/members",err)
		return models.User{}, err
	}
	
	return user, nil
}



func (data *Database) GetGroupMembers(groupID int) (ids []int, err error) {
	var rows *sql.Rows
	rows, err = data.Db.Query(`
	    SELECT DISTINCT user_id FROM (
	    SELECT sender AS user_id FROM invites WHERE group_id = ? AND status = 'accepted'
		UNION
		SELECT receiver AS user_id FROM invites WHERE group_id = ? AND status = 'accepted'
	 ) `, groupID, groupID)

    if err != nil {
	   return nil, err
	}
	defer rows.Close()

    for rows.Next() {			
		var id int																	        
	    if err = rows.Scan(&id); err != nil {																					           
	        return
	    }																													        
        ids = append(ids, id)																																					    
    }
																																								    
	if err = rows.Err(); err != nil {																																					        
		return 																																						    
	}
																																											    
	return																																						
}

