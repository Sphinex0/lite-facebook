package repository

import (
	"database/sql"
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) SaveMessage(msg *models.WSMessage) (err error) {
	args := utils.GetExecFields(msg.Message, "ID")
	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO messages
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	msg.Message.ID = int(id)
	if msg.Message.Reply != nil {
		err = data.Db.QueryRow(`
			SELECT content FROM messages WHERE id = ?
		`, msg.Message.Reply).Scan(&msg.ReplyContent)
	}

	if err != nil {
		return
	}

	// Members

	_, err = data.Db.Exec(`
		UPDATE members
		SET seen = seen + 1
		WHERE conversation_id = ?  AND member != ?
	`, msg.Message.ConversationID, msg.Message.SenderID)

	return
}

func (data *Database) GetMessagesHestories(befor, conversation_id int) (messages []models.WSMessage, err models.Error) {
	query := ` 
		SELECT M.* , COALESCE(M2.content,"") , U.id,U.first_name,U.last_name,U.nickname,U.image
		FROM messages M JOIN users U ON M.sender_id = U.id
		LEFT JOIN messages M2 ON M2.id = M.reply
		WHERE M.conversation_id = ?
		AND M.created_at < ?
		ORDER BY M.created_at DESC , M.id DESC
		LIMIT 10
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
		tab := append(utils.GetScanFields(&msg.Message), &msg.ReplyContent)
		tab = append(tab, utils.GetScanFields(&msg.UserInfo)...)
		err.Err = rows.Scan(tab...)
		if err.Err != nil {
			fmt.Println(err)
			return
		}
		messages = append(messages, msg)
	}
	fmt.Println(messages)

	return
}

func (data *Database) ReadMessages(convId int) (err error) {
	_, err = data.Db.Exec(`
		UPDATE messages
		SET seen = 1
		WHERE conversation_id = ?
	`, convId)
	if err != nil {
		fmt.Println(err)
	}
	return
}
