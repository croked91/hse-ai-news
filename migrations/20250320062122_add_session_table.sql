-- +goose Up
-- +goose StatementBegin
CREATE TYPE session_mode AS ENUM ('news', 'discussion');

-- TODO: переделать session_id на integer
CREATE TABLE session (
    id SERIAL PRIMARY KEY,
    uid VARCHAR(255) NOT NULL UNIQUE,
    mode session_mode NOT NULL DEFAULT 'news',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd
