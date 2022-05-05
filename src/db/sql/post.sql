SET search_path TO papita;

-- name: GetPost :one
SELECT * FROM post
WHERE author_id = $1 and id = $2 and is_deleted = false LIMIT 1;

-- name: ListPosts :many
SELECT * FROM post
WHERE is_deleted = false and author_id = $1
ORDER BY updated
OFFSET $2 LIMIT $3;

-- name: CreatePost :one
INSERT INTO post (
  author_id, content, updated
) VALUES (
  $1, $2, $3
)
RETURNING id, author_id, content, updated;

-- name: DeletePost :exec
UPDATE post
SET is_deleted = true
WHERE author_id = $1 and id = $2;
