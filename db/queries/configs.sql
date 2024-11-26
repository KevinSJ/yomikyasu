-- name: ListConfigs :many
SELECT * FROM configs;

-- name: CreateConfig :one
INSERT INTO configs (
    use_natural_voice, speech_speed, full_text_service_url
) VALUES (
    ?, ?, ?
) RETURNING *;

-- name: DeleteConfig :exec
DELETE FROM configs where id = ?;

-- name: UpdateConfig :exec
UPDATE configs
set use_natural_voice = ?,
speech_speed = ?,
full_text_service_url = ?
WHERE id = ? RETURNING *;
