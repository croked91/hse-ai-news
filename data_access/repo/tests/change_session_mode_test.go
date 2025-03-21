package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_ChangeSessionMode(t *testing.T) {
	tdb, err := db.New()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.MustClean(t, tdb)

	newsRepo := repo.NewNewsAI(tdb)
	type args struct {
		session domain.Session
	}
	tests := []struct {
		name    string
		args    args
		prepare func()
		want    domain.Session
		wantErr bool
	}{
		{
			name:    "should change session mode",
			args:    args{session: domain.Session{Uid: "test", Mode: "discussion"}},
			prepare: func() { assert.NoError(t, newsRepo.AddSession("test")) },
			want:    domain.Session{Uid: "test", Mode: "discussion"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}

			err := newsRepo.ChangeSessionMode(tt.args.session)
			assert.Equal(t, tt.wantErr, err != nil)
			s, err := newsRepo.GetSession(tt.args.session.Uid)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, s)
		})
	}
}
