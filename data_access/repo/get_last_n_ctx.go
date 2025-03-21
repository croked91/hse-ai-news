package repo

import "github.com/croked91/news-ai/domain"

func (r *NewsAI) GetLastNCtx(sessionID string) (domain.NLastContextList, error) {
	query := `
		SELECT id, session_id, message_type, message, sequence_number
		FROM n_last_ctx
		WHERE session_id = $1
		ORDER BY id ASC
	`
	rows, err := r.db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}

	var ctxList domain.NLastContextList
	for rows.Next() {
		var ctx domain.NLastContext
		err := rows.Scan(&ctx.ID, &ctx.SessionID, &ctx.MessageType, &ctx.Message, &ctx.SequenceNumber)
		if err != nil {
			return nil, err
		}
		ctxList = append(ctxList, ctx)
	}

	return ctxList, nil
}
