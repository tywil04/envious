package piped

import (
	"fmt"
	"strings"

	"github.com/tywil04/tubed/internal/providers/shared"
)

// Process Piped recommended videos into Tubed videos.
func processRecommendedVideos(videos []pipedRelatedStream) []shared.Video {
	var processedRecommendedVideos = make([]shared.Video, len(videos))

	for index, video := range videos {
		if video.UploaderName == "More from this channel for you" {
			continue
		}

		videoId := strings.Split(video.Url, "?v=")[1]
		authorId := strings.TrimPrefix(video.UploaderUrl, "/channel/")

		processedRecommendedVideos[index] = shared.Video{
			Title:            video.Title,
			Id:               videoId,
			ThumbnailUrl:     video.Thumbnail,
			Author:           video.UploaderName,
			AuthorId:         authorId,
			AuthorAvatarUrl:  video.UploaderAvatar,
			ShortDescription: video.ShortDescription,
			Published:        video.Uploaded,
			PublishedText:    video.UploadedDate,
			LengthSeconds:    video.Duration,
			IsShort:          video.IsShort,
			ViewCount:        video.Views,
		}
	}

	return processedRecommendedVideos
}

// Process Piped captions into tubed captions.
func processCaptions(captions []pipedSubtitle) []shared.Caption {
	var processedCaptions = make([]shared.Caption, len(captions))

	for index, caption := range captions {
		processedCaptions[index] = shared.Caption{
			Label:    caption.Name,
			Language: caption.Code,
			Url:      caption.Url,
		}
	}

	return processedCaptions
}

// Get a list of Piped instances.
// The resulting array is a map that has the keys 'display' and 'value' where 'display' is the display name of an instance and 'value' is the url.
func GetInstances() ([]map[string]string, error) {
	instancesList, err := shared.HttpGetString("https://raw.githubusercontent.com/wiki/TeamPiped/Piped-Frontend/Instances.md", nil, nil)
	if err != nil {
		return []map[string]string{}, err
	}

	splitInstances := strings.Split(instancesList, "\n")
	splitInstances = splitInstances[4 : len(splitInstances)-1] // remove first few lines of non-data and remove last \n

	var instances = make([]map[string]string, len(splitInstances))
	for index, instance := range splitInstances {
		parts := strings.Split(instance, " | ")

		name := parts[0]
		apiUrl := parts[1]
		locations := " [locations: " + parts[2] + "]"
		cors := " [cors: unknown]"

		cdn := " [cdn: %s]"
		switch parts[3] {
		case "No":
			cdn = fmt.Sprintf(cdn, "no")
		case "Yes":
			cdn = fmt.Sprintf(cdn, "yes")
		default:
			cdn = fmt.Sprintf(cdn, "unknown")
		}

		instances[index] = map[string]string{
			"display": name + locations + cdn + cors,
			"value":   apiUrl,
		}
	}

	return instances, nil
}

// Get the frontend url for the selected Piped instance.
func GetInstanceFrontend(api string) (string, error) {
	frontendUrl, err := shared.HttpGetRedirect(api, nil, nil)
	if err != nil {
		return "", err
	}
	return frontendUrl, nil
}

// Get trending videos from Piped API.
func GetTrending(instance, region string) ([]shared.Video, error) {
	decodedResponse := pipedTrendingResponse{}
	err := shared.HttpGetJson(instance+"/trending?region="+region, nil, nil, &decodedResponse)
	if err != nil {
		return []shared.Video{}, err
	}

	var trending = make([]shared.Video, len(decodedResponse))
	for index, video := range decodedResponse {
		videoId := strings.Split(video.Url, "?v=")[1]
		authorId := strings.TrimPrefix(video.UploaderUrl, "/channel/")

		trending[index] = shared.Video{
			Title:            video.Title,
			Id:               videoId,
			ThumbnailUrl:     video.Thumbnail,
			Author:           video.UploaderName,
			AuthorId:         authorId,
			AuthorAvatarUrl:  video.UploaderAvatar,
			ShortDescription: video.ShortDescription,
			Published:        video.Uploaded,
			PublishedText:    video.UploadedDate,
			LengthSeconds:    video.Duration,
			IsShort:          video.IsShort,
			ViewCount:        video.Views,
		}
	}

	return trending, nil
}

// Get video from Piped API.
func GetVideo(api, frontendUrl, videoId string) (shared.Video, error) {
	decodedResponse := pipedVideoResponse{}
	err := shared.HttpGetJson(api+"/streams/"+videoId, nil, nil, &decodedResponse)
	if err != nil {
		return shared.Video{}, err
	}

	embedUrl := frontendUrl + "/embed/" + videoId
	authorId := strings.TrimPrefix(decodedResponse.UploaderUrl, "/channel/")
	recommended := processRecommendedVideos(decodedResponse.RelatedStreams)
	captions := processCaptions(decodedResponse.Subtitles)

	newVideo := shared.Video{
		Title:             decodedResponse.Title,
		Id:                videoId,
		EmbedUrl:          embedUrl,
		ThumbnailUrl:      decodedResponse.ThumbnailUrl,
		Author:            decodedResponse.Uploader,
		AuthorId:          authorId,
		AuthorAvatarUrl:   decodedResponse.UploaderAvatar,
		Description:       decodedResponse.Description,
		Published:         decodedResponse.Uploaded,
		PublishedText:     decodedResponse.UploadedDate,
		IsLiveNow:         decodedResponse.Livestream,
		LengthSeconds:     decodedResponse.Duration,
		ViewCount:         decodedResponse.Views,
		LikeCount:         decodedResponse.Likes,
		DislikeCount:      decodedResponse.Dislikes,
		Captions:          captions,
		RecommendedVideos: recommended,
	}

	return newVideo, nil
}
