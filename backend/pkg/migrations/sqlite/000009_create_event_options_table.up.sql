-- +migrate Up
CREATE TABLE
        event_options (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                going boolean NOT NULL,
                user_id INTEGER NOT NULL,
                event_id INTEGER NOT NULL,
                FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
        );