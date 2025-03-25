-- +migrate Up
CREATE VIEW user_notifications AS
SELECT 
    n.id AS notification_id,
    n.user_id AS notified_user_id,
    n.type AS notification_type,
    IFNULL(u1.first_name, '') AS invoker_name,
    IFNULL(u1.id, 0) AS invoker_id,
    IFNULL(g.title, '') AS group_name,
    IFNULL(g.id, 0) AS group_id,
    IFNULL(e.title, '') AS event_name,
    IFNULL(e.id, 0) AS event_id,
    n.seen AS is_seen
FROM 
    notifications n
LEFT JOIN users u1 ON n.invoker_id = u1.id
LEFT JOIN groups g ON n.group_id = g.id
LEFT JOIN events e ON n.event_id = e.id;