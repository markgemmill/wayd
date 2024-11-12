package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"strconv"

	"wayd/services"
	db "wayd/services/database"

	"github.com/markgemmill/appdirs"
	"github.com/markgemmill/pathlib"
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

	appDirs := appdirs.NewAppDirs("wayd", "")
	logDir := pathlib.NewPath(appDirs.UserLogDir(), 0777)
	err := logDir.MkDirs()
	if err != nil {
		panic(err)
	}

	// timestamp := time.Now().Format("20060102-150405")
	// logFile := logDir.Join(fmt.Sprintf("wayd-%s.log.txt", timestamp))
	// logWriter, err := logFile.Open()
	// if err != nil {
	// 	panic(err)
	// }

	logWriter, err := services.LoggingSink(logDir)
	if err != nil {
		panic(err)
	}

	// defer logWriter.Close()

	logger := services.ApplicationLogger(logWriter)

	settings, err := services.NewSettings(appDirs, logger.Logger())
	if err != nil {
		logger.Fatal(fmt.Sprintf("%s", err))
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
		Logger:      logger.Logger(),
		Services: []application.Service{
			application.NewService(logger),
			application.NewService(settings),
			application.NewService(dataService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			// TODO: true or false?
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	createMenu(app, logger.Logger())

	window := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Wayd!",
		Name:      "Main",
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
		URL:              "/",
	})

	logger.Debug(fmt.Sprintf("main window id: %d", window.ID()))
	logger.Debug(fmt.Sprintf("main window name: %s", window.Name()))

	app.OnEvent("dock-window", func(e *application.CustomEvent) {
		screen, _ := window.GetScreen()
		eventData := e.Data.(map[string]any)
		dockPosition := eventData["Position"].(string)
		var x int = 0
		var y int = 0
		switch dockPosition {
		case "UR":
			x = screen.Bounds.Width - window.Width()
			y = screen.Bounds.Height - window.Height()
		case "UL":
			x = 0
			y = screen.Bounds.Height - window.Height()
		case "BR":
			x = screen.Bounds.Width - window.Width()
			y = 0
		case "BL":
			x = 0
			y = 0
		}
		window.SetPosition(x, y)
	})

	app.OnEvent("close-prompt", func(e *application.CustomEvent) {
		logger.Debug("[EVT close-prompt] started")
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

	go services.RunReminders(window, logger.Logger(), settings.Settings, reminderDelayNotice)

	logger.Debug("app.Run starting...")
	err = app.Run()
	logger.Debug("app.Run completed...")

	if err != nil {
		logger.Fatal(fmt.Sprintf("%s", err))
	}
}
