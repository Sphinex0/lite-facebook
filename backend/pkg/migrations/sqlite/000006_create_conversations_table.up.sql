-- +migrate Up
CREATE TABLE
    conversations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        entitie_one INTEGER NOT NULL,
        entitie_two_user INTEGER,
        entitie_two_group INTEGER,
        type TEXT CHECK (type IN ('private', 'group')),
        created_at INTEGER NOT NULL,
        modified_at INTEGER,
        FOREIGN KEY (entitie_one) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (entitie_two_user) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (entitie_two_group) REFERENCES groups (id) ON DELETE CASCADE,
        CHECK (
            (
                type = 'private'
                AND entitie_two_user IS NOT NULL
                AND entitie_two_group IS NULL
            )
            OR (
                type = 'group'
                AND entitie_two_group IS NOT NULL
                AND entitie_two_user IS NULL
            )
        )
    );