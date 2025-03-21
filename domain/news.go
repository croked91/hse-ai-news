package domain

import (
	"time"

	"github.com/croked91/news-ai/infrastructure/config"
)

type News struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

func (n News) Concatenate() string {
	return "Заголовок новости:" + n.Title + "\n" +
		"Текст новости:" + n.Text + "\n" +
		"Ссылка на новость:" + n.URL + "\n"
}

type NewsList []News

func (n NewsList) ToPrompt() string {
	var result string

	for _, news := range n {
		result += news.Concatenate()
	}

	prompt := config.NewsReviewPrompt()

	return prompt + result
}

type AIedNews struct {
	Content   string
	CreatedAt time.Time
}
