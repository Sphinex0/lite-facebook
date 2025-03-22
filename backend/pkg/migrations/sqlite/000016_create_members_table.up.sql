-- +migrate Up
CREATE TABLE
    members (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        member INTEGER NOT NULL,
        conversation_id INTEGER NOT NULL,
        seen INTEGER NOT NULL DEFAULT 0,
        FOREIGN KEY (member) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON DELETE CASCADE
    );