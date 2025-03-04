CREATE TABLE
    groups (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        creator INTEGER,
        title TEXT,
        descriptopn TEXT,
        createdAt DATETIME NOT NULL,
        FOREIGN KEY (creator) REFERENCES users(id) ON DELETE CASCADE,
    );