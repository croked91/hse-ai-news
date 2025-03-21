package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestGetLastNews(t *testing.T) {

	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)

	tests := []struct {
		name      string
		want      domain.AIedNews
		prepareDB func()
		errType   error
	}{
		{
			name:    "should return error if no news",
			want:    domain.AIedNews{},
			errType: repo.ErrNoNews,
		},
		{
			name: "should return news",
			want: domain.AIedNews{Content: "test"},
			prepareDB: func() {
				assert.NoError(t, newsRepo.AddNews(domain.AIedNews{Content: "test"}))
			},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareDB != nil {
				tt.prepareDB()
			}
			got, err := newsRepo.GetLastNews()
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}
