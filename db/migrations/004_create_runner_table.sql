-- +migrate Up
-- 004_create_runner_table.up.sql
CREATE TABLE IF NOT EXISTS runners (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    worker_size INTEGER NOT NULL,
    refresh_interval INTEGER NOT NULL
);

-- +migrate Down
-- 004_create_runner_table.down.sql
DROP TABLE IF EXISTS runners;

