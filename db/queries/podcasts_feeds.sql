-- name: CreatePodcastFeed :one
INSERT INTO podcast_feed(
    podcast_id,
    feed_id
) VALUES (?, ?) RETURNING *;

-- name: GetPodcastEpisodes :many
-- SELECT ;
