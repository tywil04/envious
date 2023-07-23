package tubed

import (
	"fmt"
	"strings"
)

// Process Piped recommended videos into Tubed videos.
func pipedProcessRecommendedVideos(videos []pipedRelatedStream) []Video {
	var processedRecommendedVideos = make([]Video, len(videos))

	for index, video := range videos {
		if video.UploaderName == "More from this channel for you" {
			continue
		}

		videoId := strings.Split(video.Url, "?v=")[1]
		authorId := strings.TrimPrefix(video.UploaderUrl, "/channel/")

		processedRecommendedVideos[index] = Video{
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
func pipedProcessCaptions(captions []pipedSubtitle) []Caption {
	var processedCaptions = make([]Caption, len(captions))

	for index, caption := range captions {
		processedCaptions[index] = Caption{
			Label:    caption.Name,
			Language: caption.Code,
			Url:      caption.Url,
		}
	}

	return processedCaptions
}

// Get a list of Piped instances.
// The resulting array is a map that has the keys 'display' and 'value' where 'display' is the display name of an instance and 'value' is the url.
func getPipedInstances() ([]map[string]string, error) {
	instancesList, err := httpGetString("https://raw.githubusercontent.com/wiki/TeamPiped/Piped-Frontend/Instances.md", nil, nil)
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
func getPipedInstanceFrontend(api string) (string, error) {
	frontendUrl, err := httpGetRedirect(api, nil, nil)
	if err != nil {
		return "", err
	}
	return frontendUrl, nil
}

// Get trending videos from Piped API.
func getPipedTrending(instance, region string) ([]Video, error) {
	decodedResponse := pipedTrendingResponse{}
	err := httpGetJson(instance+"/trending?region="+region, nil, nil, &decodedResponse)
	if err != nil {
		return []Video{}, err
	}

	var trending = make([]Video, len(decodedResponse))
	for index, video := range decodedResponse {
		videoId := strings.Split(video.Url, "?v=")[1]
		authorId := strings.TrimPrefix(video.UploaderUrl, "/channel/")

		trending[index] = Video{
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
func getPipedVideo(api, frontendUrl, videoId string) (Video, error) {
	decodedResponse := pipedVideoResponse{}
	err := httpGetJson(api+"/streams/"+videoId, nil, nil, &decodedResponse)
	if err != nil {
		return Video{}, err
	}

	embedUrl := frontendUrl + "/embed/" + videoId
	authorId := strings.TrimPrefix(decodedResponse.UploaderUrl, "/channel/")
	recommended := pipedProcessRecommendedVideos(decodedResponse.RelatedStreams)
	captions := pipedProcessCaptions(decodedResponse.Subtitles)

	newVideo := Video{
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
