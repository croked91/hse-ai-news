package repo

import (
	"database/sql"
	"errors"

	"github.com/croked91/news-ai/domain"
	"github.com/jackc/pgx"
)

var (
	ErrNoNews = errors.New("no news")
)

func (r *NewsAI) GetLastNews() (domain.AIedNews, error) {

	query := `
		SELECT content, created_at
		FROM news
		ORDER BY created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRow(query)

	var news domain.AIedNews
	err := row.Scan(&news.Content, &news.CreatedAt)

	return news, lastNewsError(err)

}

func lastNewsError(err error) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
		return ErrNoNews
	}
	return err
}
