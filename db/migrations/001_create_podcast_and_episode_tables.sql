-- +migrate Up
-- 001_create_podcast_and_episode_tables.up.sql
CREATE TABLE IF NOT EXISTS podcasts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    link TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS episodes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    podcast_id INTEGER NOT NULL,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    pub_date TEXT,
    file_size INTEGER,
    duration REAL,
    FOREIGN KEY (podcast_id) REFERENCES podcasts(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS episodes;
DROP TABLE IF EXISTS podcasts;
