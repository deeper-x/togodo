package main

import (
	"log"

	"github.com/deeper-x/togodo/booking"
	"github.com/deeper-x/togodo/job"
	"github.com/deeper-x/togodo/memo"
	"github.com/deeper-x/togodo/settings"
)

func main() {
	evts, err := booking.SelectEvents()
	if err != nil {
		log.Panic(err)
	}

	job.Run(settings.TickInterval, memo.Handler(evts))
}
