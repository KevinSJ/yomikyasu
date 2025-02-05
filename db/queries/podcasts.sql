-- name: ListEpisodes :many
SELECT * FROM episodes;

-- name: ListPodcasts :many
SELECT * FROM podcasts;

-- name: CreatePodcast :one
INSERT INTO podcasts(
    link,
    title,
    description
) VALUES (?, ?, ?) RETURNING *;
