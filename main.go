// main.go
package main

import (
	"embed"

	"github.com/solomonjdavid001/Dockernetes/backend/cmd"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

//go:embed config.yaml
var configFile embed.FS

func main() {
	// Pass configFile to the NewApp function
	app := cmd.NewApp(configFile)

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "Dockernetes",
		Width:             1024,
		Height:            768,
		DisableResize:     false,
		Fullscreen:        false,
		WindowStartState:  options.Maximised,
		Frameless:         false,
		MinWidth:          950,
		MinHeight:         600,
		MaxWidth:          1920,
		MaxHeight:         1080,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		AlwaysOnTop:       false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.Startup,
		OnShutdown: app.Shutdown,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Preferences: &mac.Preferences{
				TabFocusesLinks:   mac.Enabled,
				FullscreenEnabled: mac.Enabled,
			},
			About: &mac.AboutInfo{
				Title:   "Dockernetes",
				Message: "© 2025",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
