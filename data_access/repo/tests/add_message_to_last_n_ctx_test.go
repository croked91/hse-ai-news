package repo

import (
	"strconv"
	"testing"

	"github.com/croked91/news-ai/data_access/repo"
	"github.com/croked91/news-ai/domain"
	"github.com/croked91/news-ai/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestNewsAI_AddMessageToLastNCtx(t *testing.T) {
	tdb, err := db.New()
	assert.NoError(t, err)

	tRepo := repo.NewNewsAI(tdb)
	db.MustClean(t, tdb)

	tests := []struct {
		name      string
		args      domain.NLastContext
		prepareDB func()
		wantErr   bool
	}{
		{
			name: `it should be 10 messages (5 question and 5 answer) 
				in session and last question message should has sequence number 11, answer 10
				and answers and questions should be mixed
			`,
			args: domain.NLastContext{SessionID: "test-session", MessageType: "question", Message: "test-message"},
			prepareDB: func() {
				for i := range 10 {
					assert.NoError(t, tRepo.AddMessageToLastNCtx(domain.NLastContext{SessionID: "test-session", MessageType: "question", Message: "test-message" + strconv.Itoa(i)}))
					assert.NoError(t, tRepo.AddMessageToLastNCtx(domain.NLastContext{SessionID: "test-session", MessageType: "answer", Message: "test-message" + strconv.Itoa(i)}))
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareDB != nil {
				tt.prepareDB()
			}

			err := tRepo.AddMessageToLastNCtx(tt.args)
			assert.NoError(t, err)

			nLast, err := tRepo.GetLastNCtx(tt.args.SessionID)
			assert.NoError(t, err)

			assert.Len(t, nLast, 10)

			assert.Equal(t, int64(11), nLast[9].SequenceNumber)
			assert.Equal(t, nLast[9].MessageType, "question")
		})
	}
}
