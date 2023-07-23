package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	tubed := Init()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Invidious Desktop",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 27, A: 1},
		OnStartup:        tubed.Startup,
		OnShutdown:       tubed.Shutdown,
		Frameless:        true,
		Windows: &windows.Options{
			Theme: windows.Dark,
		},
		Bind: []interface{}{
			tubed,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
