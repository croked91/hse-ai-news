package repo

import (
	"database/sql"
	"errors"

	"github.com/croked91/news-ai/domain"
	"github.com/jackc/pgx"
)

var (
	ErrNoContext = errors.New("no context found")
)

func (r *NewsAI) GetCompressedContext(sessionID string) (domain.CompressedContext, error) {
	query := `
		SELECT id, session_id, ctx
		FROM compressed_ctx
		WHERE session_id = $1
	`
	row := r.db.QueryRow(query, sessionID)

	var ctx domain.CompressedContext
	err := row.Scan(&ctx.ID, &ctx.SessionID, &ctx.Context)

	return ctx, compressedContextError(err)
}

func compressedContextError(err error) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
		return ErrNoContext
	}
	return err
}
