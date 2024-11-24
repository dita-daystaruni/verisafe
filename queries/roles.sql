-- name: GetAllRoles :many
select *
from roles
limit $1
offset $2
;

-- name: GetRoleByName :one
select *
from roles
where name = $1
limit 1
;

-- name: CreateRole :one
INSERT INTO roles (name, description)
VALUES ($1, $2)
RETURNING *;


-- name: DeleteRole :exec
delete from roles
where id = $1
;


-- name: AssignRole :one
INSERT INTO user_roles (user_id, role_id, assigned_at, modified_at)
VALUES ($1, $2, NOW(),NOW())
RETURNING *;

