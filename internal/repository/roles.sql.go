// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: roles.sql

package repository

import (
	"context"

	carbon "github.com/dromara/carbon/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const assignPermissionToRole = `-- name: AssignPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW())
`

type AssignPermissionToRoleParams struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

func (q *Queries) AssignPermissionToRole(ctx context.Context, arg AssignPermissionToRoleParams) error {
	_, err := q.db.Exec(ctx, assignPermissionToRole, arg.RoleID, arg.PermissionID)
	return err
}

const assignRoleToUser = `-- name: AssignRoleToUser :exec
INSERT INTO user_roles (user_id, role_id, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW())
`

type AssignRoleToUserParams struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID int32     `json:"role_id"`
}

func (q *Queries) AssignRoleToUser(ctx context.Context, arg AssignRoleToUserParams) error {
	_, err := q.db.Exec(ctx, assignRoleToUser, arg.UserID, arg.RoleID)
	return err
}

const createPermission = `-- name: CreatePermission :one
INSERT INTO permissions (permission_name, created_at, modified_at)
VALUES ($1, NOW(), NOW())
RETURNING id, permission_name, created_at, modified_at
`

func (q *Queries) CreatePermission(ctx context.Context, permissionName string) (Permission, error) {
	row := q.db.QueryRow(ctx, createPermission, permissionName)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.PermissionName,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles (role_name, description, created_at, modified_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING id, role_name, description, created_at, modified_at
`

type CreateRoleParams struct {
	RoleName    string      `json:"role_name"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, arg.RoleName, arg.Description)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.RoleName,
		&i.Description,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePermission, id)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRole, id)
	return err
}

const getPermissionByID = `-- name: GetPermissionByID :one
SELECT id, permission_name, created_at, modified_at
FROM permissions
WHERE id = $1
`

func (q *Queries) GetPermissionByID(ctx context.Context, id int32) (Permission, error) {
	row := q.db.QueryRow(ctx, getPermissionByID, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.PermissionName,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getRoleByID = `-- name: GetRoleByID :one
SELECT id, role_name, description, created_at, modified_at
FROM roles
WHERE id = $1
`

func (q *Queries) GetRoleByID(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.RoleName,
		&i.Description,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getRoleByNameFuzzy = `-- name: GetRoleByNameFuzzy :many
SELECT id, role_name, description, created_at, modified_at
FROM roles
WHERE role_name ILIKE '%' || $1 || '%'
`

func (q *Queries) GetRoleByNameFuzzy(ctx context.Context, dollar_1 pgtype.Text) ([]Role, error) {
	rows, err := q.db.Query(ctx, getRoleByNameFuzzy, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.RoleName,
			&i.Description,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllPermissions = `-- name: ListAllPermissions :many
SELECT id, permission_name, created_at, modified_at
FROM permissions
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListAllPermissionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllPermissions(ctx context.Context, arg ListAllPermissionsParams) ([]Permission, error) {
	rows, err := q.db.Query(ctx, listAllPermissions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Permission{}
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.PermissionName,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllRoles = `-- name: ListAllRoles :many
SELECT id, role_name, description, created_at, modified_at
FROM roles
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListAllRolesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllRoles(ctx context.Context, arg ListAllRolesParams) ([]Role, error) {
	rows, err := q.db.Query(ctx, listAllRoles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.RoleName,
			&i.Description,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPermissionsForRole = `-- name: ListPermissionsForRole :many
SELECT p.id, p.permission_name, rp.created_at, rp.modified_at
FROM role_permissions rp
JOIN permissions p ON rp.permission_id = p.id
WHERE rp.role_id = $1
`

type ListPermissionsForRoleRow struct {
	ID             int32         `json:"id"`
	PermissionName string        `json:"permission_name"`
	CreatedAt      carbon.Carbon `json:"created_at"`
	ModifiedAt     carbon.Carbon `json:"modified_at"`
}

func (q *Queries) ListPermissionsForRole(ctx context.Context, roleID int32) ([]ListPermissionsForRoleRow, error) {
	rows, err := q.db.Query(ctx, listPermissionsForRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPermissionsForRoleRow{}
	for rows.Next() {
		var i ListPermissionsForRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.PermissionName,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRolesForUser = `-- name: ListRolesForUser :many
SELECT r.id, r.role_name, r.description, ur.created_at, ur.modified_at
FROM user_roles ur
JOIN roles r ON ur.role_id = r.id
WHERE ur.user_id = $1
`

type ListRolesForUserRow struct {
	ID          int32         `json:"id"`
	RoleName    string        `json:"role_name"`
	Description pgtype.Text   `json:"description"`
	CreatedAt   carbon.Carbon `json:"created_at"`
	ModifiedAt  carbon.Carbon `json:"modified_at"`
}

func (q *Queries) ListRolesForUser(ctx context.Context, userID uuid.UUID) ([]ListRolesForUserRow, error) {
	rows, err := q.db.Query(ctx, listRolesForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRolesForUserRow{}
	for rows.Next() {
		var i ListRolesForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.RoleName,
			&i.Description,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removePermissionFromRole = `-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2
`

type RemovePermissionFromRoleParams struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

func (q *Queries) RemovePermissionFromRole(ctx context.Context, arg RemovePermissionFromRoleParams) error {
	_, err := q.db.Exec(ctx, removePermissionFromRole, arg.RoleID, arg.PermissionID)
	return err
}

const removeRoleFromUser = `-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = $1 AND role_id = $2
`

type RemoveRoleFromUserParams struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID int32     `json:"role_id"`
}

func (q *Queries) RemoveRoleFromUser(ctx context.Context, arg RemoveRoleFromUserParams) error {
	_, err := q.db.Exec(ctx, removeRoleFromUser, arg.UserID, arg.RoleID)
	return err
}

const updatePermission = `-- name: UpdatePermission :one
UPDATE permissions
SET
    permission_name = COALESCE($2, permission_name),
    modified_at = NOW()
WHERE id = $1
RETURNING id, permission_name, created_at, modified_at
`

type UpdatePermissionParams struct {
	ID             int32  `json:"id"`
	PermissionName string `json:"permission_name"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error) {
	row := q.db.QueryRow(ctx, updatePermission, arg.ID, arg.PermissionName)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.PermissionName,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateRole = `-- name: UpdateRole :one
UPDATE roles
SET
    role_name = COALESCE($2, role_name),
    description = COALESCE($3, description),
    modified_at = NOW()
WHERE id = $1
RETURNING id, role_name, description, created_at, modified_at
`

type UpdateRoleParams struct {
	ID          int32       `json:"id"`
	RoleName    string      `json:"role_name"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, updateRole, arg.ID, arg.RoleName, arg.Description)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.RoleName,
		&i.Description,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
