-- +migrate Up
CREATE TABLE
    messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        conversation_id INTEGER NOT NULL,
        sender_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        seen INTEGER,
        image TEXT,
        reply INTEGER CHECK(reply != id),
        created_at INTEGER NOT NULL,
        FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON DELETE CASCADE,
        FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE
        FOREIGN KEY (reply) REFERENCES messages (id) ON DELETE CASCADE
    );
