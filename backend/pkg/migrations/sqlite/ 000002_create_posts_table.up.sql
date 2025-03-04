CREATE TABLE
    articles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email VARCHAR(255) NOT NULL UNIQUE,
        password CHAR(60) NOT NULL,
        firstName VARCHAR(255) NOT NULL,
        lastName VARCHAR(255) NOT NULL,
        datebirth DATE NOT NULL,
        avatar TEXT,
        nickname VARCHAR(255),
        aboutme VARCHAR(255),
        profileType BOOLEAN DEFAULT 0,
        createdAt DATETIME NOT NULL
    );