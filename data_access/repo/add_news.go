package repo

import "github.com/croked91/news-ai/domain"

func (r *NewsAI) AddNews(news domain.AIedNews) error {
	query := `
		INSERT INTO news (content, created_at)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, news.Content, news.CreatedAt)
	return err
}
