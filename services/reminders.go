package services

import (
	"time"
)

type ReminderService struct {
	Default   int
	NextPromt time.Time
	SyncCycle string
}

func (rs *ReminderService) Run() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			// check the time since last tick
			// does current time equal NextReminder time?
			// if so, then send reminder alert on out channel
		}
	}
}
