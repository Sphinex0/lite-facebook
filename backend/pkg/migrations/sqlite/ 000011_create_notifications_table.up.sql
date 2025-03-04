CREATE TABLE
    notifications (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        type TEXT,
        invoker_id INTEGER,
        group_id INTEGER,
        event_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
        FOREIGN KEY (invoker_id) REFERENCES users(id) ON DELETE CASCADE
        FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
    )