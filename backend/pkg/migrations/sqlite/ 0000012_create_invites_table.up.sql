-- +migrate Up
CREATE TABLE invites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    sender INTEGER NOT NULL,
    receiver INTEGER NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
    FOREIGN KEY (sender) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (receiver) REFERENCES users (id) ON DELETE CASCADE
);