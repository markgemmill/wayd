package services

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/wailsapp/wails/v3/pkg/application"
// )

// type ReminderService struct {
// 	Default   int
// 	NextPromt time.Time
// 	SyncCycle string
// }

// func (rs *ReminderService) Run() {
// 	ticker := time.NewTicker(time.Minute)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			// check the time since last tick
// 			// does current time equal NextReminder time?
// 			// if so, then send reminder alert on out channel
// 		}
// 	}
// }

// func (s *ReminderService) Name() string {
// 	return "ReminderService"
// }

// func (s *ReminderService) OnStartup(ctx context.Context, options application.ServiceOptions) error {
// 	fmt.Printf("ReminderService.OnStartup... %s\n", options.Name)
// 	return nil
// }

// func (s *ReminderService) OnShutdown() error {
// 	fmt.Println("ReminderService.OnShutdown...")
// 	return nil
// }
