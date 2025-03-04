-- +migrate Up
CREATE TABLE 
        messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender_id INTEGER,
    conversation_id INTEGER,
    content TEXT,
    created_at DATETIME NOT NULL,
    seen TEXT,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE
);
