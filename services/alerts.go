package services

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var TIMEFORMAT = "15:04:05"

// AlertTimer checks every minute if an alert needs to be raised.
// `delay“ is the initial number of minutes to delay the alert.
// `alertChannel` is an outbound channel that communicates an alert is raised.
// `delayChannel` is an inbound channel that communicates the next deley
// quantity.
func AlertTimer(delay time.Time, alertChannel chan<- time.Time, delayChannel <-chan time.Time, log *slog.Logger) {
	ticker := time.NewTicker(time.Minute)

	// this is the first alert time
	nextAlert := delay.Truncate(time.Minute)

	for {
		select {
		case <-ticker.C:
			log.Debug("[ALERTS tick]")
			currentTime := time.Now().Truncate(time.Minute)
			if currentTime.After(nextAlert) || currentTime.Equal(nextAlert) {
				// send the alert out the out channel
				log.Debug("[ALERTS tick] sending alert")
				alertChannel <- currentTime

				// wait for the next alert time
				log.Debug("[ALERTS tick] fetching next alert time")
				nextAlert = <-delayChannel
			}
		}
	}
}

var (
	SYNC_IGNORE              = "NON"
	SYNC_TO_TOP_OF_HOUR      = "TOH"
	SYNC_TO_BOTTOM_OF_HOUR   = "BOH"
	SYNC_TO_THE_HALF_HOUR    = "HLF"
	SYNC_TO_THE_QUARTER_HOUR = "QTH"
)

func DelayCalculation(currentTime time.Time, givenDelay int, syncCycle string) time.Time {
	// the next requested alert time:
	baseAlertTime := currentTime.Add(time.Duration(givenDelay) * time.Minute).Truncate(time.Minute)
	switch syncCycle {
	case SYNC_TO_THE_HALF_HOUR:
		return baseAlertTime.Round(time.Duration(30) * time.Minute)
	case SYNC_TO_THE_QUARTER_HOUR:
		return baseAlertTime.Round(time.Duration(15) * time.Minute)
	}
	return baseAlertTime
}

// RunReminders kicks off the reminder loop
// `popup` this is the gui reminder window, which pops up
// `log` the application logger
// `settings` contains the settings
// `delayNotice` is the inbound channel where the next delay in minutes
// is passed in to the alert process.
func RunReminders(popup *application.WebviewWindow, log *slog.Logger, settings *Settings, delayNotice chan int) {
	log.Debug("ALERT reminder ticker starting...")
	alertNotice := make(chan time.Time)
	calculatedDelay := make(chan time.Time)

	initialDelay := DelayCalculation(time.Now(), settings.PromptCycle, settings.SyncCycleTo)
	go AlertTimer(initialDelay, alertNotice, calculatedDelay, log)

	for {
		select {
		case t := <-alertNotice:
			log.Debug(fmt.Sprintf("ALERT %s --> alert", t.Format(TIMEFORMAT)))
			if !popup.IsVisible() {
				popup.Show()
				popup.Focus()
			}
		case d := <-delayNotice:
			log.Debug(fmt.Sprintf("ALERT next delay is %d minutes", d))
			calculatedDelay <- DelayCalculation(time.Now(), d, settings.SyncCycleTo)
		}
	}
}
