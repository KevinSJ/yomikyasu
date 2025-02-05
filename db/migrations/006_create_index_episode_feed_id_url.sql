-- 006_create_index_episode_feed_id_url.sql
-- +migrate Up
CREATE INDEX IF NOT EXISTS feed_id_url ON episodes (
    feed_id,
    url
);

-- +migrate Down
DROP INDEX IF EXISTS feed_id_url;
