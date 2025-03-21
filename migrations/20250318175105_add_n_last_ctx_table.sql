-- +goose Up
-- +goose StatementBegin
CREATE TYPE message_type AS ENUM ('question', 'answer');

CREATE TABLE n_last_ctx (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    message_type message_type NOT NULL,
    sequence_number INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd
