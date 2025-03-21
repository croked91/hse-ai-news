-- +goose Up
-- +goose StatementBegin
ALTER TABLE compressed_ctx ADD CONSTRAINT unique_session_id UNIQUE (session_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE compressed_ctx DROP CONSTRAINT unique_session_id;
-- +goose StatementEnd
