-- +migrate Up
-- 007_create_view_podcasts_episodes.sql
CREATE VIEW IF NOT EXISTS podcasts_episodes AS
SELECT pf.podcast_id, episodes.*, podcasts.link as podcast_link, podcasts.description as podcast_description, podcasts.title as podcast_title
FROM podcast_feed pf
INNER JOIN episodes ON pf.feed_id = episodes.feed_id
INNER JOIN podcasts ON pf.podcast_id = podcasts.id;

-- +migrate Down
-- 007_create_view_podcasts_episodes.sql
DROP VIEW IF EXISTS podcasts_episodes;
