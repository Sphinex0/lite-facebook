INSERT INTO users (
    nickname,
    date_birth,
    first_name,
    last_name,
    email,
    password,
    image,
    created_at
) VALUES (
    'testuser',
    952259471,
    'John',
    'Doe',
    'john.doe@example.com',
    'hashed_password123',
    'profile.jpg',
    952259471
);

INSERT INTO users (
    nickname,
    date_birth,
    first_name,
    last_name,
    email,
    password,
    image,
    created_at
) VALUES (
    'testuser2',
    952259475,
    'John2',
    'Doe2',
    'john.doe2@example.com',
    'hashed_password1234',
    'profile.jpg',
    952259475
);

INSERT INTO sessions VALUES (
    NULL,
    1,
    '550e8400-e29b-41d4-a716-446655440000',
    1977654321
)


    -- '550e8400-e29b-41d4-a716-446655440000',
    -- 1677654321,

INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (1, 1, 'new content', 'public', 1741188870, 1741188870, '', NULL, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (4, 1, 'new comment', 'public', 1741189368, 1741189368, '', NULL, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (5, 1, 'new comment', 'public', 1741189474, 1741189474, '', 4, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (6, 1, 'comment 2', 'public', 1741192462, 1741192462, '', 4, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (7, 1, 'post post post', 'public', 1741273282, 1741273282, '', NULL, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (8, 1, 'new new new post', 'public', 1741273317, 1741273317, '', NULL, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (9, 1, 'comment comment comment', 'public', 1741273435, 1741273435, '', 4, NULL);
INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (10, 1, 'comment 1', 'public', 1741273572, 1741273572, '', 8, NULL);


INSERT INTO articles (id, user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (11, 2, 'new post hahaha', 'private', 1741273282, 1741273282, '', NULL, NULL);



INSERT INTO articles (user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (2, 'This the post of dndnd', 'almost_private', 1741273572, 1741273572, '', NULL, NULL);

INSERT INTO articles (user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (2, 'This the post of dndnd 222', 'almost_private', 1741273572, 1741273572, '', NULL, NULL);

INSERT INTO articles (user_id, content, privacy, created_at, modified_at, image, parent, group_id) VALUES (2, 'mohadad', 'private', 1741273572, 1741273572, '', NULL, NULL);

INSERT INTO permited_users VALUES (NULL,14,1)