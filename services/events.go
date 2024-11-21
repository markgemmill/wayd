package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

func DockWindow(window *application.WebviewWindow) func(*application.CustomEvent) {
	return func(e *application.CustomEvent) {
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
		window.SetAlwaysOnTop(false)
	}
}
