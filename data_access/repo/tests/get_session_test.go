package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_GetSession(t *testing.T) {

	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)

	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    domain.Session
		wantErr bool
	}{
		{
			name:    "should return session",
			args:    args{uid: "test"},
			want:    domain.Session{Uid: "test", Mode: "news"},
			prepare: func() { assert.NoError(t, newsRepo.AddSession("test")) },
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}

			got, err := newsRepo.GetSession(tt.args.uid)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
