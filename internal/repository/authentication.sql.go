// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: authentication.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getUserCredentials = `-- name: GetUserCredentials :one
select user_id, username, email, password, last_login, admission_number
from login_info
where username = $1 or email = $2 or user_id = $3 or admission_number = $4
limit 1
`

type GetUserCredentialsParams struct {
	Username        string      `json:"username"`
	Email           pgtype.Text `json:"email"`
	UserID          uuid.UUID   `json:"user_id"`
	AdmissionNumber pgtype.Text `json:"admission_number"`
}

func (q *Queries) GetUserCredentials(ctx context.Context, arg GetUserCredentialsParams) (LoginInfo, error) {
	row := q.db.QueryRow(ctx, getUserCredentials,
		arg.Username,
		arg.Email,
		arg.UserID,
		arg.AdmissionNumber,
	)
	var i LoginInfo
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.LastLogin,
		&i.AdmissionNumber,
	)
	return i, err
}
