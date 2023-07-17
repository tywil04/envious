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

type Caption struct {
	Label    string `json:"label"`
	Language string `json:"language"`
	Url      string `json:"url"`
}

type Format struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Video struct {
	Title    string `json:"title"`
	Id       string `json:"id"`
	Url      string `json:"url"`
	EmbedUrl string `json:"embedUrl"`
	// DashUrl            string    `json:"dashUrl"`
	// Formats            []Format  `json:"formats"`
	ThumbnailUrl       string    `json:"thumbnailUrl"`
	Author             string    `json:"author"`
	AuthorId           string    `json:"authorId"`
	AuthorUrl          string    `json:"authorUrl"`
	AuthorThumbnailUrl string    `json:"authorThumbnailUrl"`
	Description        string    `json:"description"`
	DescriptionHtml    string    `json:"descriptionHtml"`
	Published          int       `json:"published"`
	PublishedText      string    `json:"publishedText"`
	Genre              string    `json:"genre"`
	LiveNow            bool      `json:"liveNow"`
	SubCountText       string    `json:"subCountText"`
	LengthSeconds      int       `json:"lengthSeconds"`
	AllowRatings       bool      `json:"allowRatings"`
	Rating             float32   `json:"rating"`
	IsListed           bool      `json:"isListed"`
	IsUpcoming         bool      `json:"isUpcoming"`
	ViewCount          int       `json:"viewCount"`
	ViewCountText      string    `json:"viewCountText"`
	LikeCount          int       `json:"likeCount"`
	DislikeCount       int       `json:"dislikeCount"`
	Paid               bool      `json:"paid"`
	Premium            bool      `json:"premium"`
	IsFamilyFriendly   bool      `json:"isFamilyFriendly"`
	Captions           []Caption `json:"captions"`
	RecommendedVideos  []Video   `json:"recommendedVideos"`
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
				thumbnailUrl = thumbnail.Url
				break
			}
		}

		processedResponse = append(processedResponse, Video{
			Title:         popular.Title,
			Id:            popular.VideoId,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorId:      popular.AuthorId,
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
				thumbnailUrl = thumbnail.Url
				break
			}
		}

		processedResponse = append(processedResponse, Video{
			Title:         popular.Title,
			Id:            popular.VideoId,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorId:      popular.AuthorId,
			Published:     popular.Published,
			PublishedText: popular.PublishedText,
			LiveNow:       popular.LiveNow,
			Paid:          popular.Paid,
			Premium:       popular.Premium,
		})
	}

	return processedResponse
}

func (i *InvidiousDesktop) GetVideo(videoId string) Video {
	if config.Public.SelectedInstance == "" {
		return Video{}
	}

	response, err := invidious.GetVideo(config.Public.SelectedInstance, videoId)
	if err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}

	var maxResolutionVideoThumbnailUrl string
	for _, videoThumbnail := range response.VideoThumbnails {
		if videoThumbnail.Quality == "maxres" || videoThumbnail.Quality == "maxresdefault" {
			maxResolutionVideoThumbnailUrl = videoThumbnail.Url
			break
		}
	}

	var authorThumbnailUrl string
	for _, authorThumbnail := range response.AuthorThumbnails {
		if authorThumbnail.Width == 48 && authorThumbnail.Height == 48 {
			authorThumbnailUrl = authorThumbnail.Url
			break
		}
	}

	var processedRecommendedVideos = []Video{}
	for _, recommended := range response.RecommendedVideos {
		var videoThumbnailUrl string
		for _, videoThumbnail := range recommended.VideoThumbnails {
			if videoThumbnail.Quality == "medium" {
				videoThumbnailUrl = videoThumbnail.Url
				break
			}
		}

		processedRecommendedVideos = append(processedRecommendedVideos, Video{
			Title:         recommended.Title,
			Id:            recommended.VideoId,
			ThumbnailUrl:  videoThumbnailUrl,
			Author:        recommended.Author,
			AuthorId:      recommended.AuthorId,
			AuthorUrl:     config.Public.SelectedInstance + recommended.AuthorUrl,
			LengthSeconds: recommended.LengthSeconds,
			ViewCount:     recommended.ViewCount,
			ViewCountText: recommended.ViewCountText,
		})
	}

	var processedCaptions = []Caption{}
	for _, caption := range response.Captions {
		processedCaptions = append(processedCaptions, Caption{
			Label:    caption.Label,
			Language: caption.LanguageCode,
			Url:      config.Public.SelectedInstance + caption.Url,
		})
	}

	// var processedFormats = []Format{}
	// for _, format := range response.AdaptiveFormats {
	// 	processedFormats = append(processedFormats, Format{
	// 		Type: format.Type,
	// 		Url:  format.Url,
	// 	})
	// }

	processedResponse := Video{
		Title:              response.Title,
		Id:                 response.VideoId,
		EmbedUrl:           config.Public.SelectedInstance + "/embed/" + response.VideoId,
		ThumbnailUrl:       maxResolutionVideoThumbnailUrl,
		Author:             response.Author,
		AuthorId:           response.AuthorId,
		AuthorUrl:          config.Public.SelectedInstance + response.AuthorUrl,
		AuthorThumbnailUrl: authorThumbnailUrl,
		Description:        response.Description,
		DescriptionHtml:    response.DescriptionHtml,
		Published:          response.Published,
		PublishedText:      response.PublishedText,
		Genre:              response.Genre,
		LiveNow:            response.LiveNow,
		SubCountText:       response.SubCountText,
		LengthSeconds:      response.LengthSeconds,
		AllowRatings:       response.AllowRatings,
		Rating:             response.Rating,
		IsListed:           response.IsListed,
		IsUpcoming:         response.IsUpcoming,
		ViewCount:          response.ViewCount,
		LikeCount:          response.LikeCount,
		DislikeCount:       response.DislikeCount,
		Paid:               response.Paid,
		Premium:            response.Premium,
		IsFamilyFriendly:   response.IsFamilyFriendly,
		Captions:           processedCaptions,
		RecommendedVideos:  processedRecommendedVideos,
	}

	return processedResponse
}

func (i *InvidiousDesktop) SetToken(token string) {
	config.Public.Token = token
	if err := config.Offload(); err != nil {
		runtime.LogFatal(i.ctx, err.Error())
	}
}
