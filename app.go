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

type Video struct {
	Title         string `json:"title"`
	VideoId       string `json:"videoId"`
	ThumbnailUrl  string `json:"thumbnailUrl"`
	ViewCount     int64  `json:"viewCount"`
	Author        string `json:"author"`
	AuthorUrl     string `json:"authorUrl"`
	Published     int64  `json:"published"`
	PublishedText string `json:"publishedText"`
	LiveNow       bool   `json:"liveNow"`
	Paid          bool   `json:"paid"`
	Premium       bool   `json:"premium"`
}

func (i *InvidiousDesktop) GetPopular() []Video {
	response, err := invidious.GetPopular(config.Public.SelectedInvidiousInstance)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	processedResponse := []Video{}
	for _, popular := range response {
		var thumbnailUrl string
		for _, thumbnail := range popular.VideoThumbnails {
			if thumbnail.Quality == "medium" {
				thumbnailUrl = thumbnail.URL
			}
		}

		processedResponse = append(processedResponse, Video{
			Title:         popular.Title,
			VideoId:       popular.VideoID,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorUrl:     config.Public.SelectedInvidiousInstance + popular.AuthorURL,
			Published:     popular.Published,
			PublishedText: popular.PublishedText,
			LiveNow:       popular.LiveNow,
			Paid:          popular.Paid,
			Premium:       popular.Premium,
		})
	}

	return processedResponse
}
