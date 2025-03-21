package tg

import (
	"context"
	"errors"
	"fmt"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *AINewsClient) SendNews(ctx context.Context, b *bot.Bot, update *models.Update) {
	news, err := c.newsRepo.GetLastNews()
	if errors.Is(err, repo.ErrNoNews) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Новостей нет",
		})
		return
	}

	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Произошла ошибка при обработке запроса: %v", err),
		})
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "По состоянию на" + news.CreatedAt.Format("02.01.2006") + "\n\n" + news.Content,
	})
	if err != nil {
		fmt.Println(err)
	}
}
