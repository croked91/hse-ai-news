package repo

func (r *NewsAI) AddSession(uid string) error {
	query := `
		INSERT INTO session (uid)
		VALUES ($1)
	`
	_, err := r.db.Exec(query, uid)
	return err
}
