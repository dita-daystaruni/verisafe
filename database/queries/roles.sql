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




-- name: CreatePermission :one
INSERT INTO permissions (permission_name, created_at, modified_at)
VALUES ($1, NOW(), NOW())
RETURNING *;

-- name: UpdatePermission :one
UPDATE permissions
SET
    permission_name = COALESCE($2, permission_name),
    modified_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ListAllPermissions :many
SELECT id, permission_name, created_at, modified_at
FROM permissions
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetPermissionByID :one
SELECT *
FROM permissions
WHERE id = $1;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1;


-- name: AssignPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW());

-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2;

-- name: ListPermissionsForRole :many
SELECT p.id, p.permission_name, rp.created_at, rp.modified_at
FROM role_permissions rp
JOIN permissions p ON rp.permission_id = p.id
WHERE rp.role_id = $1;

-- name: AssignRoleToUser :exec
INSERT INTO user_roles (user_id, role_id, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW());

-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = $1 AND role_id = $2;

-- name: ListRolesForUser :many
SELECT r.id, r.role_name, r.description, ur.created_at, ur.modified_at
FROM user_roles ur
JOIN roles r ON ur.role_id = r.id
WHERE ur.user_id = $1;

