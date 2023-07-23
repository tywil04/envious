package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/tywil04/tubed/internal/config"
	"github.com/tywil04/tubed/internal/providers"
	"github.com/tywil04/tubed/internal/providers/shared"
)

type Tubed struct {
	ctx context.Context

	InvidiousKnownInstances []map[string]string
	PipedKnownInstances     []map[string]string
}

func Init() *Tubed {
	t := &Tubed{}
	return t
}

func (t *Tubed) Startup(ctx context.Context) {
	if err := config.Load(); err != nil {
		runtime.LogFatal(ctx, err.Error())
	}

	t.ctx = ctx
}

func (t *Tubed) Shutdown(ctx context.Context) {
	if err := config.Offload(); err != nil {
		runtime.LogFatal(ctx, err.Error())
	}
}

func (t *Tubed) SetConfig(provider, instanceApi string) {
	frontendUrl, err := providers.GetInstanceFrontend(provider, instanceApi)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	config.Stored.Provider = provider
	config.Stored.InstanceFrontend = frontendUrl
	config.Stored.InstanceApi = instanceApi
	config.Stored.Configured = true

	err = config.Offload()
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	runtime.WindowReload(t.ctx)
}

func (t *Tubed) SetProvider(provider string) {
	config.Stored.Provider = provider
	err := config.Offload()
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}
}

func (t *Tubed) SetInstance(instanceApi string) {
	frontendUrl, err := providers.GetInstanceFrontend(config.Stored.Provider, instanceApi)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	config.Stored.InstanceFrontend = frontendUrl
	config.Stored.InstanceApi = instanceApi

	err = config.Offload()
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}
}

func (t *Tubed) GetConfigured() bool {
	return config.Stored.Configured
}

func (t *Tubed) GetProvider() string {
	return config.Stored.Provider
}

func (t *Tubed) GetInstancesApi(provider string) []map[string]string {
	switch provider {
	case "invidious":
		if t.InvidiousKnownInstances != nil {
			return t.InvidiousKnownInstances
		}
	case "piped":
		if t.PipedKnownInstances != nil {
			return t.PipedKnownInstances
		}
	}

	instances, err := providers.GetInstancesApi(provider)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	switch provider {
	case "invidious":
		t.InvidiousKnownInstances = instances
	case "piped":
		t.PipedKnownInstances = instances
	}

	return instances
}

func (t *Tubed) GetInstanceApi() string {
	return config.Stored.InstanceApi
}

func (t *Tubed) GetInstanceFrontend() string {
	return config.Stored.InstanceFrontend
}

func (t *Tubed) GetTrending() []shared.Video {
	if !config.Stored.Configured {
		return []shared.Video{}
	}

	trending, err := providers.GetTrending(config.Stored.Provider, config.Stored.InstanceApi, config.Stored.Region)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	return trending
}

func (t *Tubed) GetVideo(videoId string) shared.Video {
	if !config.Stored.Configured {
		return shared.Video{}
	}

	video, err := providers.GetVideo(config.Stored.Provider, config.Stored.InstanceApi, config.Stored.InstanceFrontend, videoId)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	return video
}
