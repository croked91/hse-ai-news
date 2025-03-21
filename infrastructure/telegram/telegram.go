package telegram

import (
	"fmt"

	"github.com/croked91/news-ai/infrastructure/config"
	"github.com/go-telegram/bot"
)

// New создает нового клиента Telegram
func New() (*bot.Bot, error) {
	token := config.TGToken()
	if token == "" {
		return nil, fmt.Errorf("переменная окружения TELEGRAM_BOT_TOKEN не установлена")
	}

	b, err := bot.New(token)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать клиента Telegram: %w", err)
	}

	return b, nil
}
