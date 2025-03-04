CREATE TABLE
    event_options (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            going TEXT,
            user_id INTEGER,
            event_id INTEGER,
            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
            FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
    )