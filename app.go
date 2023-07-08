package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"invidiousdesktop/lib/config"
	"invidiousdesktop/lib/invidious"
)

type InvidiousDesktop struct {
	ctx context.Context

	KnownInvidiousApiInstances []string
}

func NewApp() *InvidiousDesktop {
	i := &InvidiousDesktop{}
	return i
}

func (i *InvidiousDesktop) startup(ctx context.Context) {
	if err := config.Load(); err != nil {
		runtime.LogFatal(ctx, err.Error())
	}

	i.ctx = ctx
}

func (i *InvidiousDesktop) shutdown(ctx context.Context) {
	if err := config.Offload(); err != nil {
		runtime.LogFatal(ctx, err.Error())
	}
}

func (i *InvidiousDesktop) GetInvidiousApiInstances() []string {
	if i.KnownInvidiousApiInstances != nil {
		return i.KnownInvidiousApiInstances
	}

	var err error
	i.KnownInvidiousApiInstances, err = invidious.GetApiInstances()
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	return i.KnownInvidiousApiInstances
}

func (i *InvidiousDesktop) SetSelectedInvidiousInstance(instance string) bool {
	valid, err := invidious.ValidateInstance(instance)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	if valid {
		config.Public.SelectedInvidiousInstance = instance
		if err := config.Offload(); err != nil {
			runtime.LogFatal(i.ctx, err.Error())
		}
	}

	return valid
}

func (i *InvidiousDesktop) GetSelectedInvidiousInstance() string {
	return config.Public.SelectedInvidiousInstance
}
