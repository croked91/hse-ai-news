package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_AddNews(t *testing.T) {
	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)

	type args struct {
		news domain.AIedNews
	}
	tests := []struct {
		name    string
		args    args
		errType error
	}{
		{
			name:    "should add news",
			args:    args{news: domain.AIedNews{Content: "test"}},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newsRepo.AddNews(tt.args.news)
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}
