-- +migrate Up
CREATE TABLE invites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER,
    sender INTEGER,
    receiver INTEGER,
    status TEXT,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (sender) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver) REFERENCES users(id) ON DELETE CASCADE
);