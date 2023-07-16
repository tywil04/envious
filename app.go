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

func (i *InvidiousDesktop) GetApiInstances() []string {
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

func (i *InvidiousDesktop) SetSelectedInstance(instance string) bool {
	valid, err := invidious.ValidateInstance(instance)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	if valid {
		config.Public.SelectedInstance = instance
		if err := config.Offload(); err != nil {
			runtime.LogFatal(i.ctx, err.Error())
		}
	}

	return valid
}

func (i *InvidiousDesktop) GetSelectedInstance() string {
	return config.Public.SelectedInstance
}

type Video struct {
	Title         string `json:"title"`
	VideoId       string `json:"videoId"`
	ThumbnailUrl  string `json:"thumbnailUrl"`
	ViewCount     int64  `json:"viewCount"`
	Author        string `json:"author"`
	AuthorId      string `json:"authorId"`
	Published     int64  `json:"published"`
	PublishedText string `json:"publishedText"`
	LiveNow       bool   `json:"liveNow"`
	Paid          bool   `json:"paid"`
	Premium       bool   `json:"premium"`
}

func (i *InvidiousDesktop) GetPopular() []Video {
	if config.Public.SelectedInstance == "" {
		return nil
	}

	response, err := invidious.GetPopular(config.Public.SelectedInstance)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	processedResponse := []Video{}
	for _, popular := range response {
		var thumbnailUrl string
		for _, thumbnail := range popular.VideoThumbnails {
			if thumbnail.Quality == "medium" {
				thumbnailUrl = thumbnail.URL
				break
			}
		}

		processedResponse = append(processedResponse, Video{
			Title:         popular.Title,
			VideoId:       popular.VideoID,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorId:      popular.AuthorID,
			Published:     popular.Published,
			PublishedText: popular.PublishedText,
		})
	}

	return processedResponse
}

func (i *InvidiousDesktop) GetTrending() []Video {
	if config.Public.SelectedInstance == "" {
		return nil
	}

	response, err := invidious.GetTrending(config.Public.SelectedInstance)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	processedResponse := []Video{}
	for _, popular := range response {
		var thumbnailUrl string
		for _, thumbnail := range popular.VideoThumbnails {
			if thumbnail.Quality == "medium" {
				thumbnailUrl = thumbnail.URL
				break
			}
		}

		processedResponse = append(processedResponse, Video{
			Title:         popular.Title,
			VideoId:       popular.VideoID,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorId:      popular.AuthorID,
			Published:     popular.Published,
			PublishedText: popular.PublishedText,
			LiveNow:       popular.LiveNow,
			Paid:          popular.Paid,
			Premium:       popular.Premium,
		})
	}

	return processedResponse
}

func (i *InvidiousDesktop) SetToken(token string) {
	config.Public.Token = token
	if err := config.Offload(); err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}
}
