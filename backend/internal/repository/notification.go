package repository

import (
	"social-network/internal/models"
)

func (database *Database) GetUserNotifications(userID string) ([]models.Notification, error) {
	rows, err := database.Db.Query(`SELECT 
    n.id AS notification_id,
    n.type,
    u.name AS invoker_name,
    g.name AS group_name,
	n.event_id AS event_id,
FROM 
    notifications n
LEFT JOIN 
    users u ON n.invoker_id = u.id
LEFT JOIN 
    groups g ON n.group_id = g.id
WHERE 
    n.user_id = ?;
`, userID)
	if err != nil {
		return []models.Notification{}, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.ID, &notification.Type, &notification.InvokerName, &notification.GroupTitle, &notification.EventID)
		if err != nil {
			return []models.Notification{}, err
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	// count unseen notifications
	return notifications, nil
}

func (database *Database) CheckIfEventExists(EventId int) bool {
	var exists bool

	err := database.Db.QueryRow("SELECT EXISTS(SELECT title FROM events WHERE id = ?)", EventId).Scan(&exists)

	return err == nil
}

func (database *Database) InsertNotification(notification models.Notification) error {
	_, err := database.Db.Exec("INSERT INTO notifications user_id, type, invoker_id, group_id, event_id",
		notification.UserID, notification.Type, notification.InvokerID, notification.GroupID, notification.EventID)
	if err != nil {
		return err
	}
	return nil
}
