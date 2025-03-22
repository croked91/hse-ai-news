package tg

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/croked91/news-ai/domain"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	nLastMode           = "N_LAST"
	compressedMode      = "COMPRESSED"
	maxCompressedCtxLen = 120000
)

func (c *AINewsClient) Discus(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	msg := update.Message.Text
	if msg == "" {
		return
	}

	chatID := update.Message.Chat.ID

	mode := os.Getenv("NEWS_CTX_MODE")

	msg = c.getPrevContext(update.Message.Chat.ID, mode) + "\n\n" + "question:" + msg
	news, _ := c.newsRepo.GetLastNews()

	response, err := c.llm.Discus(ctx, msg+"\n"+"previous news:"+news.Content+"\n")
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Произошла ошибка при обработке запроса",
		})
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   response,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if mode == nLastMode {
		c.updateNLastCtx(chatID, msg, response)
		return
	}

	ctxToCompress := msg + "\n\n" + "ответ:" + response
	c.updateCompressedCtx(ctx, chatID, ctxToCompress)
}

func (c *AINewsClient) getPrevContext(chatID int64, mode string) string {
	// TODO: переделать session_id на integer
	sessionID := strconv.Itoa(int(chatID))

	if mode == compressedMode {
		fmt.Println("compressed mode enabled")
		compressedCtx, err := c.newsRepo.GetCompressedContext(sessionID)
		if err != nil {
			return "предыдущий контекст отсутствует"
		}

		// TODO: найти место лучше
		if utf8.RuneCountInString(compressedCtx.Context) > maxCompressedCtxLen {
			go func() {
				comprCtx, _ := c.llm.CompressCtx(context.Background(), compressedCtx.Context)
				c.newsRepo.UpsertCompressedContext(domain.CompressedContext{
					SessionID: sessionID,
					Context:   comprCtx,
				})
			}()
		}

		return compressedCtx.ToPrompt()
	}

	fmt.Println("n_last mode enabled")
	nLastCtx, err := c.newsRepo.GetLastNCtx(sessionID)
	if err != nil {
		return "предыдущий контекст отсутствует"
	}

	return nLastCtx.ToPrompt()
}

func (c *AINewsClient) updateNLastCtx(chatID int64, q, a string) {
	// TODO: сделать батчем

	// TODO: переделать session_id на integer
	sessionID := strconv.Itoa(int(chatID))

	question := domain.NLastContext{
		SessionID:   sessionID,
		Message:     q,
		MessageType: "question",
	}
	if err := c.newsRepo.AddMessageToLastNCtx(question); err != nil {
		fmt.Println(err)
	}

	answer := domain.NLastContext{
		SessionID:   sessionID,
		Message:     a,
		MessageType: "answer",
	}
	if err := c.newsRepo.AddMessageToLastNCtx(answer); err != nil {
		fmt.Println(err)
	}
}

func (c *AINewsClient) updateCompressedCtx(ctx context.Context, chatID int64, ctxToCompress string) {
	ctxAfterCompression, err := c.llm.CompressCtx(ctx, ctxToCompress)
	if err != nil {
		fmt.Println(err)
		return
	}

	newCompressedCtx := domain.CompressedContext{
		SessionID: strconv.Itoa(int(chatID)),
		Context:   ctxAfterCompression,
	}

	err = c.newsRepo.UpsertCompressedContext(newCompressedCtx)
	if err != nil {
		fmt.Println(err)
	}
}
