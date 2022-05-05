// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: author.sql

package gensql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO author (
  email, password, first_name, last_name, last_login, updated
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, email, password, first_name, last_name, last_login, created, updated
`

type CreateAuthorParams struct {
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	LastLogin sql.NullTime `db:"last_login" json:"last_login"`
	Updated   time.Time    `db:"updated" json:"updated"`
}

type CreateAuthorRow struct {
	Id        uuid.UUID    `db:"id" json:"id"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	LastLogin sql.NullTime `db:"last_login" json:"last_login"`
	Created   time.Time    `db:"created" json:"created"`
	Updated   time.Time    `db:"updated" json:"updated"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (CreateAuthorRow, error) {
	row := q.db.QueryRow(ctx, createAuthor,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.LastLogin,
		arg.Updated,
	)
	var i CreateAuthorRow
	err := row.Scan(
		&i.Id,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.LastLogin,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
UPDATE author
SET is_deleted = true
WHERE id = $1
RETURNING id, email, password, created, updated, last_login, first_name, last_name, is_deleted
`

func (q *Queries) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, email, password, created, updated, last_login, first_name, last_name, is_deleted FROM author
WHERE id = $1 and is_deleted = false LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id uuid.UUID) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.Id,
		&i.Email,
		&i.Password,
		&i.Created,
		&i.Updated,
		&i.LastLogin,
		&i.FirstName,
		&i.LastName,
		&i.IsDeleted,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, email, password, created, updated, last_login, first_name, last_name, is_deleted FROM author
WHERE is_deleted = false
ORDER BY created
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.Id,
			&i.Email,
			&i.Password,
			&i.Created,
			&i.Updated,
			&i.LastLogin,
			&i.FirstName,
			&i.LastName,
			&i.IsDeleted,
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

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE author
SET email = $2,
    password = $3,
    first_name = $4
where id = $1
RETURNING id, email, password, created, updated, last_login, first_name, last_name, is_deleted
`

type UpdateAuthorParams struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	FirstName string    `db:"first_name" json:"first_name"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRow(ctx, updateAuthor,
		arg.Id,
		arg.Email,
		arg.Password,
		arg.FirstName,
	)
	var i Author
	err := row.Scan(
		&i.Id,
		&i.Email,
		&i.Password,
		&i.Created,
		&i.Updated,
		&i.LastLogin,
		&i.FirstName,
		&i.LastName,
		&i.IsDeleted,
	)
	return i, err
}
