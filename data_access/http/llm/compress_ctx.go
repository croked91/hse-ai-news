package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/croked91/news-ai/infrastructure/config"
)

func (c *Controller) CompressCtx(
	ctx context.Context,
	ctxToCompress string,
) (string, error) {
	requestBody := map[string]interface{}{
		"model":  "deepseek-r1:32b",
		"prompt": config.NewCompressCtxPrompt() + ctxToCompress,
		"stream": false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var llmResponse LLMResponse
	err = json.Unmarshal(body, &llmResponse)
	if err != nil {
		return "", err
	}

	return llmResponse.Response, nil
}
