-- +migrate Up
CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    creator INTEGER NOT NULL,
    title TEXT NOT NULL,
    descriptopn TEXT NOT NULL,
    image TEXT,
    created_at INTEGER NOT NULL,
    FOREIGN KEY (creator) REFERENCES users(id) ON DELETE CASCADE
);