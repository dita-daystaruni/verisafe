-- name: GetUserCredentials :one
SELECT * FROM login_info
WHERE username = $1
OR email = $2
OR user_id = $3 LIMIT 1;
