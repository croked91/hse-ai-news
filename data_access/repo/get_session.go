package repo

import "github.com/croked91/news-ai/domain"

func (r *NewsAI) GetSession(uid string) (domain.Session, error) {
	query := `
		SELECT uid, mode 
		FROM session
		WHERE uid = $1
	`

	var session domain.Session
	err := r.db.QueryRow(query, uid).Scan(&session.Uid, &session.Mode)
	return session, err
}
