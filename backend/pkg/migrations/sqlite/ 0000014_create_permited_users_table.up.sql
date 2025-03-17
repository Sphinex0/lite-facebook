-- +migrate Up
CREATE TABLE
    permited_users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        article_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );