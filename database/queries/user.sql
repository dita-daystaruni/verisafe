-- name: GetUserByID :one
select users.*, sqlc.embed(userprofile), sqlc.embed(credentials)
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE id = $1
LIMIT 1
;

-- name: GetUserByEmail :one
select users.*, sqlc.embed(userprofile), sqlc.embed(credentials)
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE email = $1
LIMIT 1
;


-- name: GetUserByUsername :one
select users.*, sqlc.embed(userprofile), sqlc.embed(credentials)
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE username = $1
LIMIT 1
;


-- name: GetActiveUsers :many
select *
from users
where active = true
limit $1
offset $2
;

-- name: GetInActiveUsers :many
select *
from users
where active = false
limit $1
offset $2
;

-- name: GetAllUsers :many
select users.*, sqlc.embed(userprofile), sqlc.embed(credentials)
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
limit $1
offset $2
;
-- name: DeleteUser :exec
delete from users
where id = $1
;

-- name: CreateUser :one
INSERT INTO users (username, firstname, othernames, phone, email, gender, national_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateUserCredentials :one
INSERT INTO credentials (user_id, password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUserCredentials :one
UPDATE credentials
  SET password = $2,
  modified_at = NOW(),
  last_login = COALESCE(last_login, $3)
  WHERE user_id = $1
  RETURNING *;

-- name: CreateUserProfile :one
INSERT INTO userprofile (user_id,admission_number, bio,date_of_birth, profile_picture_url)
VALUES($1,$2,$3,$4, COALESCE(NULL, $5))
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE userprofile
  SET 
  bio = COALESCE($2, bio),
  profile_picture_url = COALESCE($3, profile_picture_url),
  modified_at = NOW()
  WHERE user_id = $1
  RETURNING *;


-- name: UpdateUserProfilePicture :one
UPDATE userprofile
  SET 
  profile_picture_url = COALESCE($2, profile_picture_url)
  WHERE user_id = $1
  RETURNING *;


-- name: GetUserProfile :one
select *
from userprofile
where user_id = $1
;

