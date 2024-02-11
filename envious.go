package main

import (
	"context"
	"slices"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/tywil04/tubed/internal/db"
	"github.com/tywil04/tubed/internal/invidious"
)

type Envious struct {
	ctx context.Context

	backendReady bool

	invidiousSession *invidious.Session
}

func Init() *Envious {
	return &Envious{}
}

func (e *Envious) startup(ctx context.Context) {
	if err := db.Read(); err != nil {
		runtime.LogFatal(ctx, "startup: "+err.Error())
	}

	if len(db.Get[[]invidious.Instance]("backend.knownInstances")) == 0 {
		knownInstances, err := invidious.GetInstances()
		if err != nil {
			runtime.LogFatal(ctx, "startup: "+err.Error())
		}
		db.Set("backend.knownInstances", knownInstances)
		go db.Write()
	}

	invidiousInstance := e.GetSelectedInstance()
	e.invidiousSession = invidious.NewSession(invidiousInstance)
	e.ctx = ctx
	e.backendReady = true
}

func (e *Envious) shutdown(ctx context.Context) {
	if err := db.Write(); err != nil {
		runtime.LogFatal(ctx, "shutdown: "+err.Error())
	}
}

// this was required to ensure everything waits until wails starts
func (e *Envious) WaitForBackend() {
	for {
		time.Sleep(time.Second)
		if e.backendReady {
			return
		}
	}
}

func (e *Envious) GetBackendConfigured() bool {
	return db.Get[bool]("backend.configured")
}

func (e *Envious) SetBackendConfigured() {
	db.Set("backend.configured", true)
}

func (e *Envious) GetInstances() []invidious.Instance {
	return db.Get[[]invidious.Instance]("backend.knownInstances")
}

func (e *Envious) SetSelectedInstance(instance invidious.Instance) {
	knownInstances := db.Get[[]invidious.Instance]("backend.knownInstances")
	index := slices.Index[[]invidious.Instance, invidious.Instance](knownInstances, instance)
	db.Set("backend.selectedInstance", index)
}

func (e *Envious) GetSelectedInstance() invidious.Instance {
	selectedInstance := db.Get[int]("backend.selectedInstance")
	knownInstances := db.Get[[]invidious.Instance]("backend.knownInstances")
	return knownInstances[selectedInstance]
}

func (e *Envious) GetTrendingVideos(trendingOptions invidious.TrendingOption) ([]invidious.Video, error) {
	return e.invidiousSession.GetTrendingVideos(trendingOptions)
}

func (e *Envious) GetVideo(videoId string) (invidious.Video, error) {
	return e.invidiousSession.GetVideo(videoId)
}

func (e *Envious) Search(query string, searchOptions invidious.SearchOption) ([]invidious.SearchItem, error) {
	return e.invidiousSession.Search(query, searchOptions)
}
