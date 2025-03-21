package repo

import (
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_AddSession(t *testing.T) {
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
		args    args
		prepare func()
		wantErr bool
	}{
		{
			name:    "should add session",
			args:    args{uid: "test"},
			wantErr: false,
		},
		{
			name:    "should return err if session already exists",
			args:    args{uid: "test"},
			prepare: func() { assert.NoError(t, newsRepo.AddSession("test")) },
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newsRepo.AddSession(tt.args.uid)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
