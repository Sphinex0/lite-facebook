-- +migrate Up
CREATE VIEW IF NOT EXISTS article_view AS
SELECT
    U.id as users_user_id,
    U.nickname,
    U.first_name,
    U.last_name,
    U.image,
    A.*,
    (
        SELECT count(like)
        FROM likes L
        WHERE
            L.article_id = A.id
            AND like = 1
    ) likes,
    (
        SELECT count(like)
        FROM likes L
        WHERE
            L.article_id = A.id
            AND like = -1
    ) dislikes,
    (
        SELECT count(*)
        FROM articles Art
        WHERE
            Art.parent = A.id
    ) comments,
    G.title,G.image
FROM articles A
    JOIN users U ON U.id = A.user_id
    LEFT JOIN groups G ON G.id = A.group_id