package config

import (
	"os"
)

type config struct {
	newsApiKey string
	appPort    string
}

var cfg = &config{}

func Init() {
	cfg.newsApiKey = os.Getenv("NEWS_API_KEY")
	// TODO: remove this
	if cfg.newsApiKey == "" {
		cfg.newsApiKey = "2aae0d5a6b2e46cea85dab566cc8a53c"
	}
	cfg.appPort = ":8080"
}

func NewsApiKey() string {
	return cfg.newsApiKey
}

func AppPort() string {
	return cfg.appPort
}

func NewsReviewPrompt() string {
	return `
		You are a professional TV host preparing a concise science news digest for the evening broadcast. Based on provided data, create a brief, lively, and informative summary. Preserve key details from headlines and news content, add smooth transitions between stories, and maintain dynamic intonation. Include source references at the end of each segment or integrate them naturally (e.g., 'More details at [website]'). Start with a greeting ('Good day! Here's the latest news...') and end with a closing phrase ('That's all for now. Stay tuned!').

		Format:

		Short paragraphs (1-3 sentences per story)
		Conversational but professional tone
		Source references at the end of each story
		Neutral but energetic delivery
		Use emoji separators (▶️) instead of bullet points
		Example structure:
		'Good day! Here are today's top stories...
		▶️ [Brief summary of first story]. Details: [link].
		▶️ [Next topic + key facts]. Source: [link].
		...
		That's all for now. Stay updated!'

		News:
	`
}

func NewsDiscusPrompt() string {
	return `
		You are a helpful assistant discussing news with users. You receive:

		News in [NEWS] format
		Previous discussion in [PREVIOUS DISCUSSION: QUESTION ANSWER] format
		Your task is to continue the discussion, considering the context.

		Guidelines:

		Start with a brief summary if [NEWS] is provided
		Maintain context using [PREVIOUS DISCUSSION]
		Ask clarifying questions
		Provide relevant context or background
		Discuss implications
		Present multiple viewpoints if applicable
		Encourage user input
		Keep responses concise but informative
		Maintain neutral, professional tone
		Avoid taking sides or making definitive predictions
		Admit when you don't know something and suggest where to find more info
		Example:
		[NEWS] Russia passed a new law about...
		[PREVIOUS DISCUSSION]
		User: What are the potential consequences?
		Assistant: Main consequences may include... Which aspect would you like to discuss further?
	`
}

func NewCompressCtxPrompt() string {
	return `
		You are a professional text compressor. Your task is to compress provided context as compactly as possible while preserving key meanings, facts, and logical connections. The compressed text should remain informative enough to support further conversation without quality loss.

		Compression rules:
		1. Remove redundant words, repetitions, and non-essential details
		2. Keep key facts, names, dates, numbers, and important events
		3. For dialogues, preserve only key lines essential for understanding
		4. Keep technical or specialized terms unchanged
		5. Ensure compressed text remains logically connected and readable

		Format:
		- Use short sentences
		- Avoid complex constructions
		- Combine multiple facts into single sentences when possible
		- Maintain chronology if important

		Example:
		Original: "Yesterday I went to the store to buy milk. The store was crowded, and I waited in line for a long time. I ended up buying milk and bread."
		Compressed: "Yesterday bought milk and bread at the store."

		Now compress the following context:
	`
}

func TGToken() string {
	return "8022229217:AAHAuql_g9zmUsVj5FKk9_O1Li5e4YBAPUY"
}
