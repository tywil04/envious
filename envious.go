package main

import (
	"context"
	"embed"
	"github.com/tywil04/tubed/internal/invidious"
	"github.com/tywil04/tubed/internal/kv"
	"github.com/tywil04/tubed/internal/proxy"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type Envious struct {
	ctx context.Context

	backendReady chan bool

	kv        *kv.DB
	invidious *invidious.Session
}

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	envious := &Envious{}

	err := wails.Run(&options.App{
		Title:  "Envious",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Middleware: proxy.Middleware,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 27, A: 1},
		OnStartup:        envious.onStartup,
		OnShutdown:       envious.onShutdown,
		Frameless:        true,
		Windows: &windows.Options{
			WebviewGpuIsDisabled: false,
			Theme:                windows.Dark,
		},
		Linux: &linux.Options{
			WebviewGpuPolicy: linux.WebviewGpuPolicyOnDemand,
		},
		Bind: []interface{}{
			envious,
		},
	})
	if err != nil {
		panic(err)
	}
}

func (e *Envious) onStartup(ctx context.Context) {
	e.ctx = ctx

	db, err := kv.NewDB("Envious" + string(os.PathSeparator) + "envious.json")
	if err != nil {
		runtime.LogFatal(ctx, err.Error())
	}
	e.kv = db

	instance, err := e.GetInvidiousInstance()
	if err != nil {
		runtime.LogFatal(ctx, err.Error())
	}
	e.invidious = invidious.NewSession(instance)
}

func (e *Envious) onShutdown(ctx context.Context) {
	err := e.kv.Close()
	if err != nil {
		runtime.LogFatal(ctx, err.Error())
	}
}

func (e *Envious) RestartApp() {
	// as close as we can get to a full backend restart
	e.onShutdown(e.ctx)
	e.onStartup(e.ctx)
	runtime.WindowReloadApp(e.ctx)
}

func (e *Envious) KVGet(key string) (string, error) {
	return e.kv.Get("frontend." + key)
}

func (e *Envious) KVSet(key string, value string) error {
	return e.kv.Set(kv.KV{Key: "frontend." + key, Value: value})
}

func (e *Envious) GetInvidiousInstances() ([]*invidious.Instance, error) {
	instances, err := invidious.GetInstances()
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (e *Envious) SetInvidiousInstance(instance *invidious.Instance) error {
	err := e.kv.SetMultiple([]kv.KV{
		{
			"invidiousInstance.ApiUrl",
			instance.ApiUrl,
		},
		{
			"invidiousInstance.Region",
			instance.Region,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *Envious) GetInvidiousInstance() (*invidious.Instance, error) {
	values, err := e.kv.GetMultiple([]string{
		"invidiousInstance.ApiUrl",
		"invidiousInstance.Region",
	})
	if err != nil {
		return nil, err
	}

	instance := &invidious.Instance{
		ApiUrl: values[0],
		Region: values[1],
	}

	return instance, nil
}

func (e *Envious) GetTrendingVideos(trendingOptions invidious.TrendingOption) ([]invidious.Video, error) {
	return e.invidious.GetTrendingVideos(trendingOptions)
}

func (e *Envious) GetVideo(videoId string) (invidious.Video, error) {
	return e.invidious.GetVideo(videoId)
}

func (e *Envious) Search(query string, searchOptions invidious.SearchOption) ([]invidious.SearchItem, error) {
	return e.invidious.Search(query, searchOptions)
}
