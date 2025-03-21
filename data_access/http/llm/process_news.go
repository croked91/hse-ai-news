package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"time"

	"github.com/croked91/news-ai/domain"
)

type LLMResponse struct {
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) ProcessNews(news domain.NewsList) {
	prompt := news.ToPrompt()

	requestBody := map[string]interface{}{
		"model":  "deepseek-r1:7b",
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var llmResponse LLMResponse
	err = json.Unmarshal(body, &llmResponse)
	if err != nil {
		fmt.Println(err)
	}

	aiedNews := domain.AIedNews{
		Content:   llmResponse.Response,
		CreatedAt: llmResponse.CreatedAt,
	}

	err = c.newsRepo.AddNews(aiedNews)
	if err != nil {
		fmt.Println(err)
	}
}
