// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: tag.sql

package gensql

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tag (
  name, updated
) VALUES (
  $1, $2
)
RETURNING id, name, created, updated
`

type CreateTagParams struct {
	Name    string    `db:"name" json:"name"`
	Updated time.Time `db:"updated" json:"updated"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, createTag, arg.Name, arg.Updated)
	var i Tag
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createTagForPost = `-- name: CreateTagForPost :one
INSERT INTO _post_tag (
  post_id, tag_id
) VALUES (
  $1, $2
)
RETURNING post_id, tag_id
`

type CreateTagForPostParams struct {
	PostID uuid.UUID `db:"post_id" json:"post_id"`
	TagID  uuid.UUID `db:"tag_id" json:"tag_id"`
}

func (q *Queries) CreateTagForPost(ctx context.Context, arg CreateTagForPostParams) (PostTag, error) {
	row := q.db.QueryRow(ctx, createTagForPost, arg.PostID, arg.TagID)
	var i PostTag
	err := row.Scan(&i.PostID, &i.TagID)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tag
where id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteTag, id)
	return err
}

const getTag = `-- name: GetTag :one
SELECT id, name, created, updated FROM tag
WHERE id = $1
`

func (q *Queries) GetTag(ctx context.Context, id uuid.UUID) (Tag, error) {
	row := q.db.QueryRow(ctx, getTag, id)
	var i Tag
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listTags = `-- name: ListTags :many
SELECT id, name, created, updated FROM tag
ORDER BY name
OFFSET $1 LIMIT $2
`

type ListTagsParams struct {
	Offset int32 `db:"offset" json:"offset"`
	Limit  int32 `db:"limit" json:"limit"`
}

func (q *Queries) ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error) {
	rows, err := q.db.Query(ctx, listTags, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Created,
			&i.Updated,
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

const listTagsForPost = `-- name: ListTagsForPost :many
select id, name from tag
inner join _post_tag
on tag.id = _post_tag.tag_id
where _post_tag.post_id = $1
`

type ListTagsForPostRow struct {
	Id   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

func (q *Queries) ListTagsForPost(ctx context.Context, postID uuid.UUID) ([]ListTagsForPostRow, error) {
	rows, err := q.db.Query(ctx, listTagsForPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTagsForPostRow
	for rows.Next() {
		var i ListTagsForPostRow
		if err := rows.Scan(&i.Id, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}