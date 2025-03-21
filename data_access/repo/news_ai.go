package repo

import "database/sql"

type NewsAI struct {
	db *sql.DB
}

func NewNewsAI(db *sql.DB) *NewsAI {
	return &NewsAI{db: db}
}
