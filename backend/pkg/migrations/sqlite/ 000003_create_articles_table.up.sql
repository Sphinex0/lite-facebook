-- +migrate Up
CREATE TABLE
    articles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        content TEXT NOT NULL,
        privacy TEXT NOT NULL,
        created_at INTEGER NOT NULL,
        modified_at INTEGER NOT NULL,
        image TEXT,
        parent INTEGER,
        group_id INTEGER,
        FOREIGN KEY (parent) REFERENCES articles (id) ON DELETE CASCADE,
        FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
    );