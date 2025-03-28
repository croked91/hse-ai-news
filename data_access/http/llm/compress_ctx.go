package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/croked91/news-ai/infrastructure/config"
)

func (c *Controller) CompressCtx(
	ctx context.Context,
	ctxToCompress string,
) (string, error) {
	requestBody := map[string]interface{}{
		"model":  "deepseek-r1:7b",
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

	idx := strings.Index(llmResponse.Response, "</think>") + len("</think>") + 1
	compress := llmResponse.Response[idx:]

	return compress, nil
}
