package repo

import (
	"github.com/croked91/news-ai/domain"
)

const (
	contextLen = 5
)

func (r *NewsAI) AddMessageToLastNCtx(nLastCtx domain.NLastContext) error {

	sessionID := nLastCtx.SessionID
	if sessionID == "" {
		sessionID = "default"
	}

	// TODO: подумать над кольцевой очередью
	// TODO: завернуть в транзакцию

	query := `
        INSERT INTO n_last_ctx (session_id, message_type, message, sequence_number)
        VALUES ($1::text, $2, $3, (
            COALESCE(
                (SELECT MAX(sequence_number) + 1
                FROM n_last_ctx
                WHERE session_id = $1::text
                AND message_type = $2),
                1
            )
        ))
    `

	_, err := r.db.Exec(query, sessionID, nLastCtx.MessageType, nLastCtx.Message)
	if err != nil {
		return err
	}

	deleteQuery := `
		DELETE FROM n_last_ctx
		WHERE session_id = $1
		AND message_type = $2
		AND sequence_number <= (
			SELECT MAX(sequence_number)
			FROM n_last_ctx
			WHERE session_id = $1
			AND message_type = $2
		) - $3
	`

	_, err = r.db.Exec(deleteQuery, sessionID, nLastCtx.MessageType, contextLen)
	return err
}
