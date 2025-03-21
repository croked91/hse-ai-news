package llm

import "github.com/croked91/news-ai/data_access/repo"

type Controller struct {
	newsRepo *repo.NewsAI
}

func NewController(apiKey string, newsRepo *repo.NewsAI) *Controller {
	return &Controller{
		newsRepo: newsRepo,
	}
}
