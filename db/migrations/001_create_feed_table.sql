-- +migrate Up
-- 001_create_feed_table.up.sql
CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    is_full_text BOOLEAN NOT NULL,
    item_since REAL,
    max_items INTEGER,
    language TEXT
);

-- +migrate Down
-- 001_create_feed_table.down.sql
DROP TABLE IF EXISTS feeds;

