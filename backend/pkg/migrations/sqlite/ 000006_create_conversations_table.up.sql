-- +migrate Up
CREATE TABLE
    conversations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        entitie_one INTEGER NOT NULL,
        entitie_two INTEGER NOT NULL,
        type TEXT,
        FOREIGN KEY (entitie_one) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (entitie_two) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (entitie_two) REFERENCES groups(id) ON DELETE CASCADE
    );