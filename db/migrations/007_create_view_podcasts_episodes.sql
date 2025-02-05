-- +migrate Up
-- 007_create_view_podcasts_episodes.sql
CREATE VIEW IF NOT EXISTS podcasts_episodes AS SELECT podcast_id, episodes.*
FROM podcast_feed pf
INNER JOIN episodes
WHERE episodes.feed_id = pf.feed_id;

-- +migrate Down
-- 007_create_view_podcasts_episodes.sql
DROP VIEW IF EXISTS podcasts_episodes;
