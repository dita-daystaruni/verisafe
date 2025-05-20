-- name: CreateReferal :one
INSERT INTO user_referals ( referring_user, referred_user, created_at, modified_at)
VALUES ( $1, $2, NOW(), NOW())
RETURNING *;
