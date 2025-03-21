-- +goose Up
-- +goose StatementBegin
CREATE TABLE compressed_ctx (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    ctx TEXT NOT NULL
);
-- +goose StatementEnd