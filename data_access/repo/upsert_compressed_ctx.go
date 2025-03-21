package repo

import (
	"errors"

	"github.com/croked91/news-ai/domain"
)

func (r *NewsAI) UpsertCompressedContext(ctx domain.CompressedContext) error {

	query := `
		INSERT INTO compressed_ctx (session_id, ctx)
		VALUES ($1, $2)
		ON CONFLICT (session_id)
		DO UPDATE SET ctx = $2
	`

	cmd, err := r.db.Exec(query, ctx.SessionID, ctx.Context)
	if err != nil {
		return err
	}

	rows, err := cmd.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
