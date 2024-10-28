package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"strconv"

	"github.com/markgemmill/wayd/services"
	db "github.com/markgemmill/wayd/services/database"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	logger := services.ApplicationLogger()

	settings, err := services.NewSettings(logger.Logger())
	if err != nil {
		log.Fatal(err)
		return
	}

	// reminderDelayMinutes := settings.Settings.PromptCycle
	reminderDelayNotice := make(chan int)

	dataService, err := db.NewDatabaseService(settings.DatabasePath(), logger.Logger())
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "wayd",
		Description: "What Are You Doing?",
		Services: []application.Service{
			application.NewService(settings),
			application.NewService(logger),
			application.NewService(dataService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Wayd!",
		Width:     390,
		Height:    844,
		MinWidth:  390,
		MinHeight: 844,
		MaxWidth:  300,
		MaxHeight: 844,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(0, 0, 0),
		URL:              "/new",
	})

	popup := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Name:          "name",
		Title:         "name",
		Width:         390,
		Height:        390,
		MinWidth:      390,
		MinHeight:     390,
		MaxWidth:      390,
		MaxHeight:     390,
		AlwaysOnTop:   true,
		URL:           "/prompt",
		Hidden:        true,
		DisableResize: true,
		Centered:      true,
		Frameless:     true,
	})

	app.OnEvent("close-prompt", func(e *application.CustomEvent) {
		logger.Debug("[EVT close-prompt] started")
		popup.Hide()
		eventData := e.Data.(map[string]any)

		logger.Debug(fmt.Sprintf("[EVT close-prompt] Close Prompt Data: %v", eventData))
		delay := eventData["Delay"].(string)

		overrideDelayMinutes, err := strconv.ParseInt(delay, 0, 0)
		if err != nil {
			reminderDelayNotice <- settings.Settings.PromptCycle
			return
		}

		logger.Debug(fmt.Sprintf("[EVT close-prompt] Setting delay minutes: %d", overrideDelayMinutes))
		reminderDelayNotice <- int(overrideDelayMinutes)
		logger.Debug("[EVT close-prompt] completed")
	})

	app.OnEvent("new-entry", func(e *application.CustomEvent) {
		// reset the delay to the default when starting a new
		// timer entry.
		logger.Debug("[EVT new-entry] started")
		reminderDelayNotice <- settings.Settings.PromptCycle
		logger.Debug("[EVT new-entry] completed")
	})

	go services.RunReminders(popup, logger.Logger(), settings.Settings, reminderDelayNotice)

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
