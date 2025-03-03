// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package repository

import (
	"context"

	carbon "github.com/dromara/carbon/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, firstname, othernames, phone, email, gender, national_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, username, firstname, othernames, phone, email, gender, active, national_id, created_at, modified_at
`

type CreateUserParams struct {
	Username   string      `json:"username"`
	Firstname  string      `json:"firstname"`
	Othernames string      `json:"othernames"`
	Phone      pgtype.Text `json:"phone"`
	Email      pgtype.Text `json:"email"`
	Gender     pgtype.Text `json:"gender"`
	NationalID pgtype.Text `json:"national_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Firstname,
		arg.Othernames,
		arg.Phone,
		arg.Email,
		arg.Gender,
		arg.NationalID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Othernames,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Active,
		&i.NationalID,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const createUserCredentials = `-- name: CreateUserCredentials :one
INSERT INTO credentials (user_id, password)
VALUES ($1, $2)
RETURNING user_id, password, last_login, created_at, modified_at
`

type CreateUserCredentialsParams struct {
	UserID   uuid.UUID `json:"user_id"`
	Password *string   `json:"password"`
}

func (q *Queries) CreateUserCredentials(ctx context.Context, arg CreateUserCredentialsParams) (Credential, error) {
	row := q.db.QueryRow(ctx, createUserCredentials, arg.UserID, arg.Password)
	var i Credential
	err := row.Scan(
		&i.UserID,
		&i.Password,
		&i.LastLogin,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO userprofile (user_id,admission_number, bio,date_of_birth, profile_picture_url)
VALUES($1,$2,$3,$4, COALESCE(NULL, $5))
RETURNING user_id, admission_number, bio, vibe_points, date_of_birth, profile_picture_url, campus, last_seen, created_at, modified_at
`

type CreateUserProfileParams struct {
	UserID          uuid.UUID          `json:"user_id"`
	AdmissionNumber pgtype.Text        `json:"admission_number"`
	Bio             pgtype.Text        `json:"bio"`
	DateOfBirth     pgtype.Timestamptz `json:"date_of_birth"`
	Column5         interface{}        `json:"column_5"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, createUserProfile,
		arg.UserID,
		arg.AdmissionNumber,
		arg.Bio,
		arg.DateOfBirth,
		arg.Column5,
	)
	var i Userprofile
	err := row.Scan(
		&i.UserID,
		&i.AdmissionNumber,
		&i.Bio,
		&i.VibePoints,
		&i.DateOfBirth,
		&i.ProfilePictureUrl,
		&i.Campus,
		&i.LastSeen,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users
where id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getActiveUsers = `-- name: GetActiveUsers :many
select id, username, firstname, othernames, phone, email, gender, active, national_id, created_at, modified_at
from users
where active = true
limit $1
offset $2
`

type GetActiveUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetActiveUsers(ctx context.Context, arg GetActiveUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getActiveUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Firstname,
			&i.Othernames,
			&i.Phone,
			&i.Email,
			&i.Gender,
			&i.Active,
			&i.NationalID,
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

const getAllUsers = `-- name: GetAllUsers :many
select users.id, users.username, users.firstname, users.othernames, users.phone, users.email, users.gender, users.active, users.national_id, users.created_at, users.modified_at, userprofile.user_id, userprofile.admission_number, userprofile.bio, userprofile.vibe_points, userprofile.date_of_birth, userprofile.profile_picture_url, userprofile.campus, userprofile.last_seen, userprofile.created_at, userprofile.modified_at, credentials.user_id, credentials.password, credentials.last_login, credentials.created_at, credentials.modified_at
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
limit $1
offset $2
`

type GetAllUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllUsersRow struct {
	ID          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Firstname   string        `json:"firstname"`
	Othernames  string        `json:"othernames"`
	Phone       pgtype.Text   `json:"phone"`
	Email       pgtype.Text   `json:"email"`
	Gender      pgtype.Text   `json:"gender"`
	Active      pgtype.Bool   `json:"active"`
	NationalID  pgtype.Text   `json:"national_id"`
	CreatedAt   carbon.Carbon `json:"created_at"`
	ModifiedAt  carbon.Carbon `json:"modified_at"`
	Userprofile Userprofile   `json:"userprofile"`
	Credential  Credential    `json:"credential"`
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllUsersRow{}
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Firstname,
			&i.Othernames,
			&i.Phone,
			&i.Email,
			&i.Gender,
			&i.Active,
			&i.NationalID,
			&i.CreatedAt,
			&i.ModifiedAt,
			&i.Userprofile.UserID,
			&i.Userprofile.AdmissionNumber,
			&i.Userprofile.Bio,
			&i.Userprofile.VibePoints,
			&i.Userprofile.DateOfBirth,
			&i.Userprofile.ProfilePictureUrl,
			&i.Userprofile.Campus,
			&i.Userprofile.LastSeen,
			&i.Userprofile.CreatedAt,
			&i.Userprofile.ModifiedAt,
			&i.Credential.UserID,
			&i.Credential.Password,
			&i.Credential.LastLogin,
			&i.Credential.CreatedAt,
			&i.Credential.ModifiedAt,
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

const getInActiveUsers = `-- name: GetInActiveUsers :many
select id, username, firstname, othernames, phone, email, gender, active, national_id, created_at, modified_at
from users
where active = false
limit $1
offset $2
`

type GetInActiveUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetInActiveUsers(ctx context.Context, arg GetInActiveUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getInActiveUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Firstname,
			&i.Othernames,
			&i.Phone,
			&i.Email,
			&i.Gender,
			&i.Active,
			&i.NationalID,
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

const getUserByEmail = `-- name: GetUserByEmail :one
select users.id, users.username, users.firstname, users.othernames, users.phone, users.email, users.gender, users.active, users.national_id, users.created_at, users.modified_at, userprofile.user_id, userprofile.admission_number, userprofile.bio, userprofile.vibe_points, userprofile.date_of_birth, userprofile.profile_picture_url, userprofile.campus, userprofile.last_seen, userprofile.created_at, userprofile.modified_at, credentials.user_id, credentials.password, credentials.last_login, credentials.created_at, credentials.modified_at
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE email = $1
LIMIT 1
`

type GetUserByEmailRow struct {
	ID          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Firstname   string        `json:"firstname"`
	Othernames  string        `json:"othernames"`
	Phone       pgtype.Text   `json:"phone"`
	Email       pgtype.Text   `json:"email"`
	Gender      pgtype.Text   `json:"gender"`
	Active      pgtype.Bool   `json:"active"`
	NationalID  pgtype.Text   `json:"national_id"`
	CreatedAt   carbon.Carbon `json:"created_at"`
	ModifiedAt  carbon.Carbon `json:"modified_at"`
	Userprofile Userprofile   `json:"userprofile"`
	Credential  Credential    `json:"credential"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Othernames,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Active,
		&i.NationalID,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.Userprofile.UserID,
		&i.Userprofile.AdmissionNumber,
		&i.Userprofile.Bio,
		&i.Userprofile.VibePoints,
		&i.Userprofile.DateOfBirth,
		&i.Userprofile.ProfilePictureUrl,
		&i.Userprofile.Campus,
		&i.Userprofile.LastSeen,
		&i.Userprofile.CreatedAt,
		&i.Userprofile.ModifiedAt,
		&i.Credential.UserID,
		&i.Credential.Password,
		&i.Credential.LastLogin,
		&i.Credential.CreatedAt,
		&i.Credential.ModifiedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select users.id, users.username, users.firstname, users.othernames, users.phone, users.email, users.gender, users.active, users.national_id, users.created_at, users.modified_at, userprofile.user_id, userprofile.admission_number, userprofile.bio, userprofile.vibe_points, userprofile.date_of_birth, userprofile.profile_picture_url, userprofile.campus, userprofile.last_seen, userprofile.created_at, userprofile.modified_at, credentials.user_id, credentials.password, credentials.last_login, credentials.created_at, credentials.modified_at
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE id = $1
LIMIT 1
`

type GetUserByIDRow struct {
	ID          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Firstname   string        `json:"firstname"`
	Othernames  string        `json:"othernames"`
	Phone       pgtype.Text   `json:"phone"`
	Email       pgtype.Text   `json:"email"`
	Gender      pgtype.Text   `json:"gender"`
	Active      pgtype.Bool   `json:"active"`
	NationalID  pgtype.Text   `json:"national_id"`
	CreatedAt   carbon.Carbon `json:"created_at"`
	ModifiedAt  carbon.Carbon `json:"modified_at"`
	Userprofile Userprofile   `json:"userprofile"`
	Credential  Credential    `json:"credential"`
}

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Othernames,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Active,
		&i.NationalID,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.Userprofile.UserID,
		&i.Userprofile.AdmissionNumber,
		&i.Userprofile.Bio,
		&i.Userprofile.VibePoints,
		&i.Userprofile.DateOfBirth,
		&i.Userprofile.ProfilePictureUrl,
		&i.Userprofile.Campus,
		&i.Userprofile.LastSeen,
		&i.Userprofile.CreatedAt,
		&i.Userprofile.ModifiedAt,
		&i.Credential.UserID,
		&i.Credential.Password,
		&i.Credential.LastLogin,
		&i.Credential.CreatedAt,
		&i.Credential.ModifiedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select users.id, users.username, users.firstname, users.othernames, users.phone, users.email, users.gender, users.active, users.national_id, users.created_at, users.modified_at, userprofile.user_id, userprofile.admission_number, userprofile.bio, userprofile.vibe_points, userprofile.date_of_birth, userprofile.profile_picture_url, userprofile.campus, userprofile.last_seen, userprofile.created_at, userprofile.modified_at, credentials.user_id, credentials.password, credentials.last_login, credentials.created_at, credentials.modified_at
from users
LEFT JOIN userprofile on userprofile.user_id = users.id
LEFT JOIN credentials on credentials.user_id = users.id
WHERE username = $1
LIMIT 1
`

type GetUserByUsernameRow struct {
	ID          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Firstname   string        `json:"firstname"`
	Othernames  string        `json:"othernames"`
	Phone       pgtype.Text   `json:"phone"`
	Email       pgtype.Text   `json:"email"`
	Gender      pgtype.Text   `json:"gender"`
	Active      pgtype.Bool   `json:"active"`
	NationalID  pgtype.Text   `json:"national_id"`
	CreatedAt   carbon.Carbon `json:"created_at"`
	ModifiedAt  carbon.Carbon `json:"modified_at"`
	Userprofile Userprofile   `json:"userprofile"`
	Credential  Credential    `json:"credential"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Othernames,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Active,
		&i.NationalID,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.Userprofile.UserID,
		&i.Userprofile.AdmissionNumber,
		&i.Userprofile.Bio,
		&i.Userprofile.VibePoints,
		&i.Userprofile.DateOfBirth,
		&i.Userprofile.ProfilePictureUrl,
		&i.Userprofile.Campus,
		&i.Userprofile.LastSeen,
		&i.Userprofile.CreatedAt,
		&i.Userprofile.ModifiedAt,
		&i.Credential.UserID,
		&i.Credential.Password,
		&i.Credential.LastLogin,
		&i.Credential.CreatedAt,
		&i.Credential.ModifiedAt,
	)
	return i, err
}

const getUserProfile = `-- name: GetUserProfile :one
select user_id, admission_number, bio, vibe_points, date_of_birth, profile_picture_url, campus, last_seen, created_at, modified_at
from userprofile
where user_id = $1
`

func (q *Queries) GetUserProfile(ctx context.Context, userID uuid.UUID) (Userprofile, error) {
	row := q.db.QueryRow(ctx, getUserProfile, userID)
	var i Userprofile
	err := row.Scan(
		&i.UserID,
		&i.AdmissionNumber,
		&i.Bio,
		&i.VibePoints,
		&i.DateOfBirth,
		&i.ProfilePictureUrl,
		&i.Campus,
		&i.LastSeen,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateUserCredentials = `-- name: UpdateUserCredentials :one
UPDATE credentials
  SET password = $2,
  modified_at = NOW(),
  last_login = COALESCE(last_login, $3)
  WHERE user_id = $1
  RETURNING user_id, password, last_login, created_at, modified_at
`

type UpdateUserCredentialsParams struct {
	UserID    uuid.UUID     `json:"user_id"`
	Password  *string       `json:"password"`
	LastLogin carbon.Carbon `json:"last_login"`
}

func (q *Queries) UpdateUserCredentials(ctx context.Context, arg UpdateUserCredentialsParams) (Credential, error) {
	row := q.db.QueryRow(ctx, updateUserCredentials, arg.UserID, arg.Password, arg.LastLogin)
	var i Credential
	err := row.Scan(
		&i.UserID,
		&i.Password,
		&i.LastLogin,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE userprofile
  SET 
  bio = COALESCE($2, bio),
  profile_picture_url = COALESCE($3, profile_picture_url),
  modified_at = NOW()
  WHERE user_id = $1
  RETURNING user_id, admission_number, bio, vibe_points, date_of_birth, profile_picture_url, campus, last_seen, created_at, modified_at
`

type UpdateUserProfileParams struct {
	UserID            uuid.UUID   `json:"user_id"`
	Bio               pgtype.Text `json:"bio"`
	ProfilePictureUrl pgtype.Text `json:"profile_picture_url"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, updateUserProfile, arg.UserID, arg.Bio, arg.ProfilePictureUrl)
	var i Userprofile
	err := row.Scan(
		&i.UserID,
		&i.AdmissionNumber,
		&i.Bio,
		&i.VibePoints,
		&i.DateOfBirth,
		&i.ProfilePictureUrl,
		&i.Campus,
		&i.LastSeen,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateUserProfilePicture = `-- name: UpdateUserProfilePicture :one
UPDATE userprofile
  SET 
  profile_picture_url = COALESCE($2, profile_picture_url)
  WHERE user_id = $1
  RETURNING user_id, admission_number, bio, vibe_points, date_of_birth, profile_picture_url, campus, last_seen, created_at, modified_at
`

type UpdateUserProfilePictureParams struct {
	UserID            uuid.UUID   `json:"user_id"`
	ProfilePictureUrl pgtype.Text `json:"profile_picture_url"`
}

func (q *Queries) UpdateUserProfilePicture(ctx context.Context, arg UpdateUserProfilePictureParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, updateUserProfilePicture, arg.UserID, arg.ProfilePictureUrl)
	var i Userprofile
	err := row.Scan(
		&i.UserID,
		&i.AdmissionNumber,
		&i.Bio,
		&i.VibePoints,
		&i.DateOfBirth,
		&i.ProfilePictureUrl,
		&i.Campus,
		&i.LastSeen,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
