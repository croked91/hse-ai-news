package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func MustRegisterNewJob(f any, period time.Duration) (cancel func()) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	j, err := s.NewJob(
		gocron.DurationJob(period),
		gocron.NewTask(f),
	)
	if err != nil {
		panic(err)
	}

	s.Start()

	fmt.Println("крона с ID:", j.ID(), "запущена")

	return func() {
		err := s.Shutdown()
		if err != nil {
			log.Default().Printf("failed to shutdown scheduler: %v", err)
		}
	}
}
