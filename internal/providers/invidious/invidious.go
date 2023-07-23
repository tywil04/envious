package invidious

import (
	"fmt"
	"strings"

	"github.com/tywil04/tubed/internal/providers/shared"
)

// Process Invidious video thumbnails into a single, sensible thumbnail url.
func processVideoThumbnailUrl(thumbnails []invidiousVideoThumbnail) string {
	var processedThumbnailUrl string

	for _, videoThumbnail := range thumbnails {
		if videoThumbnail.Quality == "medium" {
			processedThumbnailUrl = videoThumbnail.Url
			break
		}
	}

	return processedThumbnailUrl
}

// Process Invidious avatar thumbnails into a single, sensible thumbnail url.
func processAuthorAvatarUrl(avatars []invidiousAuthorThumbnail) string {
	var processedAvatarUrl string

	for _, avatar := range avatars {
		if avatar.Width == 48 && avatar.Height == 48 {
			processedAvatarUrl = avatar.Url
			break
		}
	}

	return processedAvatarUrl
}

// Process Invidious recommended videos into Tubed videos.
func processRecommendedVideos(videos []invidiousRecommendedVideo) []shared.Video {
	var processedRecommendedVideos = make([]shared.Video, len(videos))

	for index, video := range videos {
		videoThumbnail := processVideoThumbnailUrl(video.VideoThumbnails)

		processedRecommendedVideos[index] = shared.Video{
			Title:         video.Title,
			Id:            video.VideoId,
			ThumbnailUrl:  videoThumbnail,
			Author:        video.Author,
			AuthorId:      video.AuthorId,
			LengthSeconds: video.LengthSeconds,
			ViewCount:     video.ViewCount,
			ViewCountText: video.ViewCountText,
		}
	}

	return processedRecommendedVideos
}

// Process Invidious captions into Tubed captions.
func processCaptions(api string, captions []invidiousCaption) []shared.Caption {
	var processedCaptions = make([]shared.Caption, len(captions))

	for index, caption := range captions {
		processedCaptions[index] = shared.Caption{
			Label:    caption.Label,
			Language: caption.LanguageCode,
			Url:      api + caption.Url,
		}
	}

	return processedCaptions
}

// Get a list of Invidious instances.
// The resulting array is a map that has the keys 'display' and 'value' where 'display' is the display name of an instance and 'value' is the url.
func GetInstances() ([]map[string]string, error) {
	decodedResponse := [][]any{}
	err := shared.HttpGetJson("https://api.invidious.io/instances.json?pretty=1&sort_by=type,users,api", nil, nil, &decodedResponse)
	if err != nil {
		return []map[string]string{}, err
	}

	var instances = make([]map[string]string, len(decodedResponse))
	for index, instance := range decodedResponse {
		mapInstance := instance[1].(map[string]any)

		if mapInstance["api"] != nil && mapInstance["api"].(bool) && mapInstance["type"].(string) == "https" {
			apiUrl := mapInstance["uri"].(string)
			name := strings.TrimPrefix(apiUrl, mapInstance["type"].(string)+"://")
			locations := " [locations: " + mapInstance["region"].(string) + "]"
			cdn := " [cdn: unknown]"

			cors := " [cors: %s]"
			switch mapInstance["cors"].(bool) {
			case false:
				cors = fmt.Sprintf(cors, "no")
			case true:
				cors = fmt.Sprintf(cors, "yes")
			default:
				cors = fmt.Sprintf(cors, "unknown")
			}

			instances[index] = map[string]string{
				"display": name + locations + cdn + cors,
				"value":   apiUrl,
			}
		} else {
			// this will ALWAYS be the first instance without an api or that is not using HTTPS because in
			// the request I am asking for the result to be sorted by wether or not its api is enabled and
			// by type

			instances = instances[:index] // make array fit values because nothing else is going to be added
			break
		}
	}

	return instances, nil
}

// Get the frontend url for the selected Invidious instance.
func GetInstanceFrontend(api string) (string, error) {
	return api, nil
}

// Get trending videos from Invidious API.
func GetTrending(api, region string) ([]shared.Video, error) {
	decodedResponse := invidiousTrendingResponse{}
	err := shared.HttpGetJson(api+"/api/v1/trending?region="+region, nil, nil, &decodedResponse)
	if err != nil {
		return []shared.Video{}, err
	}

	var trending = make([]shared.Video, len(decodedResponse))
	for index, video := range decodedResponse {
		thumbnail := processVideoThumbnailUrl(video.VideoThumbnails)

		trending[index] = shared.Video{
			Title:         video.Title,
			Id:            video.VideoId,
			ThumbnailUrl:  thumbnail,
			ViewCount:     video.ViewCount,
			Author:        video.Author,
			AuthorId:      video.AuthorId,
			Published:     video.Published,
			PublishedText: video.PublishedText,
			IsLiveNow:     video.LiveNow,
			IsPaid:        video.Paid,
			IsPremium:     video.Premium,
		}
	}

	return trending, nil
}

// Get video from Invidious API.
func GetVideo(api, frontendUrl, videoId string) (shared.Video, error) {
	decodedResponse := invidiousVideoResponse{}
	err := shared.HttpGetJson(api+"/api/v1/videos/"+videoId, nil, nil, &decodedResponse)
	if err != nil {
		return shared.Video{}, err
	}

	embedUrl := frontendUrl + "/embed/" + decodedResponse.VideoId
	videoThumbnail := processVideoThumbnailUrl(decodedResponse.VideoThumbnails)
	authorAvatar := processAuthorAvatarUrl(decodedResponse.AuthorThumbnails)
	recommended := processRecommendedVideos(decodedResponse.RecommendedVideos)
	captions := processCaptions(api, decodedResponse.Captions)

	newVideo := shared.Video{
		Title:             decodedResponse.Title,
		Id:                decodedResponse.VideoId,
		EmbedUrl:          embedUrl,
		ThumbnailUrl:      videoThumbnail,
		Author:            decodedResponse.Author,
		AuthorId:          decodedResponse.AuthorId,
		AuthorAvatarUrl:   authorAvatar,
		Description:       decodedResponse.DescriptionHtml,
		Published:         decodedResponse.Published,
		PublishedText:     decodedResponse.PublishedText,
		Genre:             decodedResponse.Genre,
		IsLiveNow:         decodedResponse.LiveNow,
		SubCountText:      decodedResponse.SubCountText,
		LengthSeconds:     decodedResponse.LengthSeconds,
		AllowRatings:      decodedResponse.AllowRatings,
		Rating:            decodedResponse.Rating,
		IsListed:          decodedResponse.IsListed,
		IsUpcoming:        decodedResponse.IsUpcoming,
		ViewCount:         decodedResponse.ViewCount,
		LikeCount:         decodedResponse.LikeCount,
		DislikeCount:      decodedResponse.DislikeCount,
		IsPaid:            decodedResponse.Paid,
		IsPremium:         decodedResponse.Premium,
		IsFamilyFriendly:  decodedResponse.IsFamilyFriendly,
		Captions:          captions,
		RecommendedVideos: recommended,
	}

	return newVideo, nil
}
