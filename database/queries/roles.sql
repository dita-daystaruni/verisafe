-- name: CreateRole :one
INSERT INTO roles (role_name, description, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING *;


-- name: UpdateRole :one
UPDATE roles
SET
    role_name = COALESCE($2, role_name),
    description = COALESCE($3, description),
    modified_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ListAllRoles :many
SELECT id, role_name, description, created_at, modified_at
FROM roles
ORDER BY id
LIMIT $1 OFFSET $2;


-- name: GetRoleByID :one
SELECT *
FROM roles
WHERE id = $1;

-- name: GetRoleByNameFuzzy :many
SELECT *
FROM roles
WHERE role_name ILIKE '%' || $1 || '%';

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;


