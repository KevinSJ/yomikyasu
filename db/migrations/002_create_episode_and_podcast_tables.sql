-- 002_create_podcast_and_episode_tables.sql
-- +migrate Up
CREATE TABLE IF NOT EXISTS podcasts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    link TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS episodes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid text NOT NULL,
    feed_id INTEGER NOT NULL,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    pub_date TEXT,
    file_size REAL,
    duration REAL,
    audio_content BLOB,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS episodes;
DROP TABLE IF EXISTS podcasts;
