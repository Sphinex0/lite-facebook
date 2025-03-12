package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (data *Database) CreateConversation(conv *models.Conversation) (err error) {
	args := utils.GetExecFields(conv, "ID")

	res, err := data.Db.Exec(fmt.Sprintf(`
		INSERT INTO conversations
		VALUES (NULL, %v) 
	`, utils.Placeholders(len(args))),
		args...)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	conv.ID = int(id)

	return
}

func (data *Database) VerifyConversation(id1, id2 int, type_obj string) (err error) {
	param := `WHERE entitie_two_group = ?`
	if type_obj == "private" {
		param = `WHERE entitie_two_user = ?`
	}
	var result int
	err = data.Db.QueryRow(fmt.Sprintf(`
		SELECT id
		FROM conversations
		WHERE entitie_one = ? AND %v`,
		param),
		id1, id2).Scan(&result)
	return
}

func (data *Database) GetConversations(id int) (conversations []models.ConversationsInfo, err error) {
	query := `
		SELECT *
		FROM conversations
		WHERE entitie_one = ? OR entitie_two_user = ?
	`
	rows, err := data.Db.Query(query, id, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var conv models.ConversationsInfo
		err1 := rows.Scan(utils.GetScanFields(&conv.Conversation)...)
		if err1 != nil {
			fmt.Println(err1)
		}
		conversations = append(conversations, conv)
	}

	for i, conv := range conversations {
		if conv.Conversation.Type == "group" {
			row := data.GetGroupById(*conv.Conversation.Entitie_two_group)
			err1 := row.Scan(utils.GetScanFields(&conversations[i].Group)...)
			if err1 != nil {
				fmt.Println(err1)
			}
		} else {
			var err1 error
			conversations[i].UserInfo, err1 = data.GetUserByID(*conv.Conversation.Entitie_two_user)
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	}

	return
}
