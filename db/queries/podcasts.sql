-- name: ListEpisodes :many
SELECT * FROM episodes;

-- name: CreateEpisode :one
INSERT INTO episodes (
    uuid,
    feed_id,
    url,
    title,
    description,
    pub_date,
    file_size,
    duration,
    audio_content
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: ListPodcasts :many
SELECT * FROM podcasts;

-- name: CreatePodcast :one
INSERT INTO podcasts(
    link,
    title,
    description
) VALUES (?, ?, ?) RETURNING *;
