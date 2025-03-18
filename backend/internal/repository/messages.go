package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveMessage(msg *models.Message) (err error) {
	args := utils.GetExecFields(msg, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO messages
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	msg.ID = int(id)
	return
}

func (data *Database) GetMessagesHestories(befor, conversation_id int) (messages []models.WSMessage, err models.Error) {
	query := ` 
		SELECT M.* , U.id,U.first_name,U.last_name,U.nickname,U.image
		FROM messages M JOIN users U ON M.sender_id = U.id
		WHERE M.conversation_id = ?
		AND M.created_at < ?
		ORDER BY M.created_at , M.id
	`
	var rows *sql.Rows
	rows, err.Err = data.Db.Query(query, conversation_id, befor)
	if err.Err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	fmt.Println("hh")
	for rows.Next() {
		var msg models.WSMessage
		tab := utils.GetScanFields(&msg.Message)
		tab = append(tab, utils.GetScanFields(&msg.UserInfo)...)
		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			fmt.Println(err)
			return
		}
		messages = append(messages, msg)
	}
	// fmt.Println(messages)

	return
}
