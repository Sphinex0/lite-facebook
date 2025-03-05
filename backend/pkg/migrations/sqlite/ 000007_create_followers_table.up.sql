-- +migrate Up
CREATE TABLE followers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    follower INTEGER NOT NULL,
    status TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower) REFERENCES users(id) ON DELETE CASCADE
);