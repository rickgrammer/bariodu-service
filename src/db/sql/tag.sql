SET search_path TO papita;

-- name: GetTag :one
SELECT * FROM tag
WHERE id = $1;

-- name: ListTags :many
SELECT * FROM tag
ORDER BY name
OFFSET $1 LIMIT $2;

-- name: ListTagsForPost :many
select id, name from tag
inner join _post_tag
on tag.id = _post_tag.tag_id
where _post_tag.post_id = $1;

-- name: CreateTag :one
INSERT INTO tag (
  name, updated
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateTagForPost :one
INSERT INTO _post_tag (
  post_id, tag_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tag
where id = $1;
