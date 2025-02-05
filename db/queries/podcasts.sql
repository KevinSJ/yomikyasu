-- name: ListEpisodes :many
SELECT * FROM episodes;

-- name: GetEpisodeExistsByUrlAndFeedId :one
SELECT EXISTS (SELECT 1 FROM episodes WHERE url = ? and feed_id = ? LIMIT 1);

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
