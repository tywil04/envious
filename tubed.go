package main

import (
	"context"
	"slices"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/tywil04/tubed/internal/db"
	"github.com/tywil04/tubed/internal/invidious"
)

type Tubed struct {
	ctx context.Context
}

func Init() *Tubed {
	t := &Tubed{}
	return t
}

func (t *Tubed) Startup(ctx context.Context) {
	if err := db.Read(); err != nil {
		runtime.LogFatal(ctx, "startup: "+err.Error())
	}

	if len(db.Get[[]invidious.Instance]("backend.knownInstances")) == 0 {
		knownInstances, err := invidious.GetInstances()
		if err != nil {
			runtime.LogFatal(ctx, "startup: "+err.Error())
		}
		db.Set("backend.knownInstances", knownInstances)
	}

	t.ctx = ctx
}

func (t *Tubed) Shutdown(ctx context.Context) {
	if err := db.Write(); err != nil {
		runtime.LogFatal(ctx, "shutdown: "+err.Error())
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
	selectedInstanceIndex := slices.Index[[]invidious.Instance, invidious.Instance](knownInstances, instance)
	db.Set("backend.selectedInstance", selectedInstanceIndex)
}

func (t *Tubed) GetSelectedInstance() invidious.Instance {
	knownInstances := db.Get[[]invidious.Instance]("backend.knownInstances")
	selectedInstanceIndex := db.Get[int]("backend.selectedInstance")
	return knownInstances[selectedInstanceIndex]
}

func (t *Tubed) GetTrendingVideos() []invidious.Video {
	videos, err := invidious.GetTrendingVideos(t.GetSelectedInstance())
	if err != nil {
		runtime.LogFatal(t.ctx, "get trending videos: "+err.Error())
	}
	return videos
}

func (t *Tubed) GetPopularVideos() []invidious.Video {
	videos, err := invidious.GetPopularVideos(t.GetSelectedInstance())
	if err != nil {
		runtime.LogFatal(t.ctx, "get popular videos: "+err.Error())
	}
	return videos
}

func (t *Tubed) GetVideo(videoId string) invidious.Video {
	video, err := invidious.GetVideo(t.GetSelectedInstance(), videoId)
	if err != nil {
		runtime.LogFatalf(t.ctx, "get video [%s]: "+err.Error(), videoId)
	}
	return video
}

func (t *Tubed) Search(query string, searchOptions invidious.SearchOptions) []invidious.SearchItem {
	searchItems, err := invidious.Search(t.GetSelectedInstance(), query, searchOptions)
	if err != nil {
		runtime.LogFatalf(t.ctx, "search [q: %s, searchOptions: %v]: "+err.Error(), query, searchOptions)
	}
	return searchItems
}
