package cron

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex

func TestMustRegisterNewJob(t *testing.T) {
	type args struct {
		f      func() (<-chan struct{}, func())
		period time.Duration
	}
	tests := []struct {
		name       string
		args       args
		wantCancel func()
	}{
		{
			name: "should write to channel 2 times",
			args: args{
				f: func() (<-chan struct{}, func()) {
					ch := make(chan struct{})
					var count int

					return ch, func() {
						ch <- struct{}{}

						mu.Lock()
						defer mu.Unlock()

						count++
						fmt.Println("count:", count)
						if count == 2 {
							close(ch)
						}
					}
				},
				period: 1 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch, f := tt.args.f()

			cancel := MustRegisterNewJob(f, tt.args.period)
			t.Cleanup(cancel)

			for range ch {
				fmt.Println("foo")
			}
		})
	}
}
