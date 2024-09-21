-- +migrate Up
CREATE TABLE IF NOT EXISTS podcast_feed (
    podcast_id INTEGER NOT NULL,
    feed_id INTEGER NOT NULL,
    PRIMARY KEY (podcast_id, feed_id),
    FOREIGN KEY (podcast_id) REFERENCES podcasts(id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS podcast_feed;

