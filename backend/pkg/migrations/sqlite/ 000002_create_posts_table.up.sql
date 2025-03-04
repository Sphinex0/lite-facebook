CREATE TABLE articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    content TEXT,
    privacy TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    modified_at DATETIME,
    image TEXT,
    parent INTEGER,
    group_id INTEGER,
    FOREIGN KEY (parent) REFERENCES articles(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
