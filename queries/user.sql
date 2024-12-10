-- name: GetUserByID :one
select *
from users
where id = $1
limit 1
;

-- name: GetUserByEmail :one
select *
from users
where email = $1
limit 1
;

-- name: GetUserByUsername :one
select *
from users
where username = $1
limit 1
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
select *
from users
limit $1
offset $2
;

-- name: DeleteUser :exec
delete from users
where id = $1
;
commit
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
INSERT INTO userprofile (user_id,admission_number, bio,campus,date_of_birth, profile_picture_url)
VALUES($1,$2,$3,$4,$5,'no-profile')
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE userprofile
  SET 
  admission_number = COALESCE($2, admission_number),
  bio = COALESCE($3, bio),
  profile_picture_url = COALESCE($4, profile_picture_url),
  campus = COALESCE($5, campus),
  modified_at = NOW()
  WHERE user_id = $1
  RETURNING *;

