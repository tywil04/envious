package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/tywil04/tubed/internal/db"
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
	if err := db.Read(); err != nil {
		runtime.LogFatal(ctx, "startup: "+err.Error())
	}

	t.ctx = ctx
}

func (t *Tubed) Shutdown(ctx context.Context) {
	if err := db.Write(); err != nil {
		runtime.LogFatal(ctx, "shutdown: "+err.Error())
	}
}

func (t *Tubed) DBSet(key string, value any) {
	db.Set(key, value)
}

func (t *Tubed) DBGet(key string, returnType string) any {
	switch returnType {
	case "bool":
		return db.Get[bool](key)
	case "string":
		return db.Get[string](key)
	case "int":
		return db.Get[int](key)
	case "uint":
		return db.Get[uint](key)
	case "byte":
		return db.Get[byte](key)
	case "float32":
		return db.Get[float32](key)
	case "[]bool":
		return db.Get[[]bool](key)
	case "[]string":
		return db.Get[[]string](key)
	case "[]int":
		return db.Get[[]int](key)
	case "[]uint":
		return db.Get[[]uint](key)
	case "[]byte":
		return db.Get[[]byte](key)
	case "[]float32":
		return db.Get[[]float32](key)
	default:
		return db.Get[any](key)
	}
}

func (t *Tubed) DBGetWithDefault(key string, returnType string, defaultValue any) any {
	switch returnType {
	case "bool":
		return db.GetWithDefault[bool](key, defaultValue)
	case "string":
		return db.GetWithDefault[string](key, defaultValue)
	case "int":
		return db.GetWithDefault[int](key, defaultValue)
	case "uint":
		return db.GetWithDefault[uint](key, defaultValue)
	case "byte":
		return db.GetWithDefault[byte](key, defaultValue)
	case "[]bool":
		return db.GetWithDefault[[]bool](key, defaultValue)
	case "[]string":
		return db.GetWithDefault[[]string](key, defaultValue)
	case "[]int":
		return db.GetWithDefault[[]int](key, defaultValue)
	case "[]uint":
		return db.GetWithDefault[[]uint](key, defaultValue)
	case "[]byte":
		return db.GetWithDefault[[]byte](key, defaultValue)
	default:
		return db.GetWithDefault[any](key, defaultValue)
	}
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

func (t *Tubed) GetTrending() []shared.Video {
	if !db.Get[bool]("backend.configured") {
		return []shared.Video{}
	}

	trending, err := providers.GetTrending(db.Get[string]("backend.provider"), db.Get[string]("backend.instanceApi"), db.Get[string]("backend.region"))
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	return trending
}

func (t *Tubed) GetVideo(videoId string) shared.Video {
	if !db.Get[bool]("backend.configured") {
		return shared.Video{}
	}

	video, err := providers.GetVideo(db.Get[string]("backend.provider"), db.Get[string]("backend.instanceApi"), db.Get[string]("backend.instanceFrontend"), videoId)
	if err != nil {
		runtime.LogFatal(t.ctx, err.Error())
	}

	return video
}
