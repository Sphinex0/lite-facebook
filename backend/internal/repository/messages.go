package repository

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

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
			row := data.GetGroupById(conv.Conversation.Entitie_two_group)
			err1 := row.Scan(utils.GetScanFields(&conversations[i].Group)...)
			if err1 != nil {
				fmt.Println(err1)
			}
		} else {
			var err1 error
			conversations[i].UserInfo, err1 = data.GetUserByID(conv.Conversation.Entitie_two_group)
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	}

	return
}
