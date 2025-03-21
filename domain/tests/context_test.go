package domain

import (
	"testing"

	"github.com/croked91/news-ai/domain"
	"github.com/stretchr/testify/assert"
)

func TestNLastContextList_Concatenate(t *testing.T) {
	tests := []struct {
		name string
		n    domain.NLastContextList
		want string
	}{
		{
			name: "should concatenate messages",
			n: domain.NLastContextList{
				domain.NLastContext{MessageType: "test", Message: "test"},
				domain.NLastContext{MessageType: "test2", Message: "test2"},
			},
			want: "Твой предыдущий диалог с пользователем:\ntest: test\ntest2: test2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.n.Concatenate()
			assert.Equal(t, tt.want, got)
		})
	}
}
