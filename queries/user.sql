-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetActiveUsers :many
SELECT * FROM users
WHERE active = true LIMIT $1 OFFSET $2;

-- name: GetInActiveUsers :many
SELECT * FROM users
WHERE active = false LIMIT $1 OFFSET $2;

-- name: GetAllUsers :many
SELECT * FROM users
LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
