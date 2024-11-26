-- name: ListFeeds :many
SELECT * FROM feeds;

-- name: CreateFeed :one
INSERT INTO feeds (
    url, is_full_text, item_since, max_items, language
) VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds where id = ?;

-- name: UpdateFeed :exec
UPDATE feeds
set url = ?,
is_full_text = ?,
item_since = ?,
max_items = ?,
language = ?
WHERE id = ? RETURNING *;
