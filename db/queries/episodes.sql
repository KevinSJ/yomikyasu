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


-- name: GetEpisodeExistsByUrlAndFeedId :one
SELECT EXISTS (SELECT 1 FROM episodes WHERE url = ? and feed_id = ? LIMIT 1);

-- name: GetEpisodeContentByUuid :one
SELECT audio_content from episodes where uuid = ?
