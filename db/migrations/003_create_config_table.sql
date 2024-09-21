-- +migrate Up
-- 003_create_config_table.up.sql
CREATE TABLE IF NOT EXISTS configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    use_natural_voice BOOLEAN NOT NULL,
    speech_speed REAL NOT NULL,
    full_text_service_url TEXT,
    refresh_interval INTEGER
);


-- +migrate Down
-- 003_create_config_table.down.sql
DROP TABLE IF EXISTS configs;
