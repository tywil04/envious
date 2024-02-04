package main

import (
	"embed"

	"github.com/tywil04/tubed/internal/proxy"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	tubed := Init()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Tubed",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Middleware: proxy.Middleware,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 27, A: 1},
		OnDomReady:       tubed.startup,
		OnShutdown:       tubed.shutdown,
		Frameless:        true,
		Windows: &windows.Options{
			WebviewGpuIsDisabled: false,
			Theme:                windows.Dark,
		},
		Linux: &linux.Options{
			WebviewGpuPolicy: linux.WebviewGpuPolicyOnDemand,
		},
		Bind: []interface{}{
			tubed,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
