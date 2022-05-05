SET search_path TO papita;

-- name: GetAuthor :one
SELECT * FROM author
WHERE id = $1 and is_deleted = false LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM author
WHERE is_deleted = false
ORDER BY created;

-- name: CreateAuthor :one
INSERT INTO author (
  email, password, first_name, last_name, last_login, updated
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, email, password, first_name, last_name, last_login, created, updated;

-- name: UpdateAuthor :one
UPDATE author
SET email = $2,
    password = $3,
    first_name = $4
where id = $1
RETURNING *;

-- name: DeleteAuthor :exec
UPDATE author
SET is_deleted = true
WHERE id = $1
RETURNING *;
