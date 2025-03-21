package repo

import "github.com/croked91/news-ai/domain"

func (r *NewsAI) ChangeSessionMode(session domain.Session) error {
	query := `
		UPDATE session
		SET mode = $2
		WHERE uid = $1
	`
	_, err := r.db.Exec(query, session.Uid, session.Mode)
	return err
}
