package tg

import (
	"context"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/go-telegram/bot"
)

type LLM interface {
	Discus(ctx context.Context, question string) (answer string, err error)
	CompressCtx(ctx context.Context, ctxToCompress string) (ctxAfterCompression string, err error)
}

type AINewsClient struct {
	newsBot  *bot.Bot
	newsRepo *repo.NewsAI
	llm      LLM
}

func NewAINewsClient(
	b *bot.Bot,
	r *repo.NewsAI,
	l LLM,
) *AINewsClient {

	c := AINewsClient{
		newsBot:  b,
		newsRepo: r,
		llm:      l,
	}

	c.newsBot.RegisterHandler(bot.HandlerTypeMessageText, "/news", bot.MatchTypeExact, c.SendNews)
	c.newsBot.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, c.Discus)

	return &c
}

func (c *AINewsClient) Start(ctx context.Context) {
	c.newsBot.Start(ctx)
}

// Close
func (c *AINewsClient) Close(ctx context.Context) {
	_, _ = c.newsBot.Close(ctx)
}
