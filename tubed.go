package main

import (
	"context"
	"slices"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/tywil04/tubed/internal/db"
	"github.com/tywil04/tubed/internal/invidious"
)

type Tubed struct {
	ctx context.Context

	backendReady bool

	invidiousSession *invidious.Session
}

func Init() *Tubed {
	t := &Tubed{}
	return t
}

func (t *Tubed) startup(ctx context.Context) {
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

	invidiousInstance := t.GetSelectedInstance()
	t.invidiousSession = invidious.NewSession(invidiousInstance)
	t.ctx = ctx
	t.backendReady = true
}

func (t *Tubed) shutdown(ctx context.Context) {
	if err := db.Write(); err != nil {
		runtime.LogFatal(ctx, "shutdown: "+err.Error())
	}
}

// this was required to ensure everything waits until wails starts
func (t *Tubed) WaitForBackend() {
	for {
		time.Sleep(time.Second)
		if t.backendReady {
			return
		}
	}
}

func (t *Tubed) GetBackendConfigured() bool {
	return db.Get[bool]("backend.configured")
}

func (t *Tubed) SetBackendConfigured() {
	db.Set("backend.configured", true)
}

func (t *Tubed) GetInstances() []invidious.Instance {
	return db.Get[[]invidious.Instance]("backend.knownInstances")
}

func (t *Tubed) SetSelectedInstance(instance invidious.Instance) {
	knownInstances := db.Get[[]invidious.Instance]("backend.knownInstances")
	index := slices.Index[[]invidious.Instance, invidious.Instance](knownInstances, instance)
	db.Set("backend.selectedInstance", index)
}

func (t *Tubed) GetSelectedInstance() invidious.Instance {
	selectedInstance := db.Get[int]("backend.selectedInstance")
	knownInstances := db.Get[[]invidious.Instance]("backend.knownInstances")
	return knownInstances[selectedInstance]
}

func (t *Tubed) GetTrendingVideos(trendingOptions invidious.TrendingOption) ([]invidious.Video, error) {
	return t.invidiousSession.GetTrendingVideos(trendingOptions)
}

func (t *Tubed) GetVideo(videoId string) (invidious.Video, error) {
	return t.invidiousSession.GetVideo(videoId)
}

func (t *Tubed) Search(query string, searchOptions invidious.SearchOption) ([]invidious.SearchItem, error) {
	return t.invidiousSession.Search(query, searchOptions)
}
