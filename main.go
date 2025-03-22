package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/croked91/news-ai/data_access/http/llm"
	"github.com/croked91/news-ai/data_access/http/news_data"
	"github.com/croked91/news-ai/data_access/http/tg"
	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/infrastructure/config"
	"github.com/croked91/news-ai/infrastructure/cron"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/croked91/news-ai/infrastructure/telegram"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config.Init()

	apiKey := config.NewsApiKey()

	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	newsRepo := repo.NewNewsAI(d)
	llmController := llm.NewController(apiKey, newsRepo)
	newsController := news_data.NewController(apiKey, llmController)

	fmt.Println("mode:", os.Getenv("NEWS_CTX_MODE"))

	cron.MustRegisterNewJob(newsController.ScienceNewsS, 1*time.Hour)

	b, err := telegram.New()
	if err != nil {
		log.Fatal(err)
	}

	newsBot := tg.NewAINewsClient(b, newsRepo, llmController)
	defer newsBot.Close(ctx)
	newsBot.Start(ctx)
}
