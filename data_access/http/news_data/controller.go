package news_data

import (
	"github.com/croked91/news-ai/domain"
)

type llm interface {
	ProcessNews(news domain.NewsList)
}

type Controller struct {
	apiKey string
	llm    llm
}

func NewController(apiKey string, l llm) *Controller {
	return &Controller{
		apiKey: apiKey,
		llm:    l,
	}
}
