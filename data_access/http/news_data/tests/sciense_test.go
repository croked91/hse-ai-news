package news_data

import (
	"testing"

	"github.com/croked91/news-ai/data_access/http/news_data"
	"github.com/croked91/news-ai/domain"
)

type llmMock struct{}

func (l *llmMock) ProcessNews(news domain.NewsList) {}

func TestController_ScienceNewsS(t *testing.T) {
	lm := &llmMock{}
	type fields struct {
		apiKey string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "should get science news",
			fields: fields{apiKey: "2aae0d5a6b2e46cea85dab566cc8a53c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := news_data.NewController(tt.fields.apiKey, lm)
			c.ScienceNewsS()
		})
	}
}
