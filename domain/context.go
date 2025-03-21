package domain

import "github.com/croked91/news-ai/infrastructure/config"

type CompressedContext struct {
	ID        int64  `json:"id"`
	SessionID string `json:"session_id"`
	Context   string `json:"context"`
}

func (c CompressedContext) ToPrompt() string {
	return config.NewsDiscusPrompt() + "Новости и предыдущий диалог с пользователем:\n" + c.Context
}

type NLastContext struct {
	ID             int64  `json:"id"`
	SessionID      string `json:"session_id"`
	Message        string `json:"message"`
	MessageType    string `json:"message_type"`
	SequenceNumber int64  `json:"sequence_number"`
}

type NLastContextList []NLastContext

func (n NLastContextList) Concatenate() string {
	var result string

	for _, ctx := range n {
		result += ctx.MessageType + ": " + ctx.Message + "\n"
	}

	result = "Новости и предыдущий диалог с пользователем:\n" + result

	return result
}

func (n NLastContextList) ToPrompt() string {
	return config.NewsDiscusPrompt() + n.Concatenate()
}
