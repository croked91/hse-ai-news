package domain

import (
	"testing"

	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/config"
)

func TestNews_Concatenate(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Link        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should concatenate news",
			fields: fields{
				Title:       "test",
				Description: "test",
				Link:        "test",
			},
			want: "Заголовок новости:test\n" +
				"Описание новости:test\n" +
				"Ссылка на новость:test\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := domain.News{
				Title: tt.fields.Title,
				Text:  tt.fields.Description,
				URL:   tt.fields.Link,
			}
			if got := n.Concatenate(); got != tt.want {
				t.Errorf("News.Concatenate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewsList_ToPrompt(t *testing.T) {
	prompt := config.NewsReviewPrompt()
	tests := []struct {
		name string
		n    domain.NewsList
		want string
	}{
		{
			name: "should concatenate news",
			n: domain.NewsList{
				domain.News{Title: "test", Text: "test", URL: "test"},
				domain.News{Title: "test2", Text: "test2", URL: "test2"},
			},
			want: prompt + "\n\n\nЗаголовок новости:test\n" +
				"Описание новости:test\n" +
				"Ссылка на новость:test\n" +
				"\n\n\nЗаголовок новости:test2\n" +
				"Описание новости:test2\n" +
				"Ссылка на новость:test2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ToPrompt(); got != tt.want {
				t.Errorf("NewsList.ToPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
