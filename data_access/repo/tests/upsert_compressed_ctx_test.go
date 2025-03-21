package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_UpsertCompressedContext(t *testing.T) {
	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)

	type args struct {
		ctx domain.CompressedContext
	}
	tests := []struct {
		name      string
		args      args
		errType   error
		prepareDB func()
	}{
		{
			name: "it should create new compressed context",
			args: args{ctx: domain.CompressedContext{SessionID: "test-session", Context: "test context"}},
		},
		{
			name: "it should update existing compressed context",
			prepareDB: func() {
				assert.NoError(t, newsRepo.UpsertCompressedContext(
					domain.CompressedContext{SessionID: "test-session2", Context: "test context"},
				))
			},
			args: args{ctx: domain.CompressedContext{SessionID: "test-session2", Context: "updated context"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareDB != nil {
				tt.prepareDB()
			}

			err := newsRepo.UpsertCompressedContext(tt.args.ctx)
			assert.ErrorIs(t, err, tt.errType)

			var savedContext string
			err = tdb.QueryRow("SELECT ctx FROM compressed_ctx WHERE session_id = $1", tt.args.ctx.SessionID).Scan(&savedContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.args.ctx.Context, savedContext)
		})
	}
}
