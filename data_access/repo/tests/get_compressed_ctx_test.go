package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_GetCompressedContext(t *testing.T) {
	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)

	type args struct {
		sessionID string
	}
	tests := []struct {
		name      string
		args      args
		want      domain.CompressedContext
		errType   error
		prepareDB func()
	}{
		{
			name:    "it should return error if no context",
			args:    args{sessionID: "nonexistent-session"},
			want:    domain.CompressedContext{},
			errType: repo.ErrNoContext,
		},
		{
			name: "it should return context",
			args: args{sessionID: "test-session"},
			want: domain.CompressedContext{SessionID: "test-session", Context: "test context"},
			prepareDB: func() {
				assert.NoError(t, newsRepo.UpsertCompressedContext(
					domain.CompressedContext{SessionID: "test-session", Context: "test context"},
				))
			},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareDB != nil {
				tt.prepareDB()
			}
			got, err := newsRepo.GetCompressedContext(tt.args.sessionID)
			assert.Equal(t, tt.want.Context, got.Context)
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}
