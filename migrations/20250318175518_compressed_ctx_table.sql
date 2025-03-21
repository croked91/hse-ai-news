-- +goose Up
-- +goose StatementBegin
CREATE TABLE compressed_ctx (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL, -- TODO: переделать session_id на integer
    ctx TEXT NOT NULL,
);
-- +goose StatementEnd