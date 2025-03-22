package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Controller) Discus(
	ctx context.Context,
	discus string,
) (string, error) {

	fmt.Println("Прилетел вопрос:", discus)

	requestBody := map[string]interface{}{
		"model":  "deepseek-r1:7b",
		"prompt": discus,
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
	answer := llmResponse.Response[idx:]

	fmt.Println("Прилетел ответ:", answer)

	return answer, nil
}
