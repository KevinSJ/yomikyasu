-- name: GetPodcastEpisodesByPodcastId :many
SELECT * FROM podcasts_episodes where podcast_id = ?;
