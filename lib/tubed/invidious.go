package tubed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getInvidiousInstances() ([]map[string]string, error) {
	response, err := http.Get("https://api.invidious.io/instances.json?pretty=1&sort_by=type,users")
	if err != nil {
		return []map[string]string{}, err
	}

	decoded := [][]any{}
	json.NewDecoder(response.Body).Decode(&decoded)

	instances := []map[string]string{}
	for _, instance := range decoded {
		mapInstance := instance[1].(map[string]any)

		if mapInstance["api"] != nil && mapInstance["api"].(bool) && mapInstance["type"].(string) == "https" {
			instanceLocations := " [locations: " + mapInstance["region"].(string) + "]"

			instanceCdn := " [cdn: unknown]"

			instanceCors := " [cors: %s]"
			switch mapInstance["cors"].(bool) {
			case false:
				instanceCors = fmt.Sprintf(instanceCors, "no")
			case true:
				instanceCors = fmt.Sprintf(instanceCors, "yes")
			default:
				instanceCors = fmt.Sprintf(instanceCors, "unknown")
			}

			apiUrl := mapInstance["uri"].(string)

			instanceName := strings.TrimPrefix(apiUrl, mapInstance["type"].(string)+"://")

			instances = append(instances, map[string]string{
				"display": instanceName + instanceLocations + instanceCdn + instanceCors,
				"value":   apiUrl,
			})
		}
	}

	return instances, nil
}

func getInvidiousInstanceFrontend(api string) (string, error) {
	return api, nil
}

type invidiousTrendingResponse []struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	VideoId         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		Url     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	LengthSeconds   int64  `json:"lengthSeconds"`
	ViewCount       int64  `json:"viewCount"`
	Author          string `json:"author"`
	AuthorId        string `json:"authorId"`
	AuthorUrl       string `json:"authorUrl"`
	Published       int64  `json:"published"`
	PublishedText   string `json:"publishedText"`
	Description     string `json:"description"`
	DescriptionHtml string `json:"descriptionHtml"`
	LiveNow         bool   `json:"liveNow"`
	Paid            bool   `json:"paid"`
	Premium         bool   `json:"premium"`
}

func getInvidiousTrending(instance, region string) ([]Video, error) {
	response, err := http.Get(instance + "/api/v1/trending?region=" + region)
	if err != nil {
		return nil, err
	}

	decodedResponse := invidiousTrendingResponse{}
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	trending := []Video{}

	for _, popular := range decodedResponse {
		var thumbnailUrl string
		for _, thumbnail := range popular.VideoThumbnails {
			if thumbnail.Quality == "medium" {
				thumbnailUrl = thumbnail.Url
				break
			}
		}

		trending = append(trending, Video{
			Title:         popular.Title,
			Id:            popular.VideoId,
			ThumbnailUrl:  thumbnailUrl,
			ViewCount:     popular.ViewCount,
			Author:        popular.Author,
			AuthorId:      popular.AuthorId,
			Published:     popular.Published,
			PublishedText: popular.PublishedText,
			IsLiveNow:     popular.LiveNow,
			IsPaid:        popular.Paid,
			IsPremium:     popular.Premium,
		})
	}

	return trending, nil
}

type invidiousVideoResponse struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	VideoId         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		Url     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	Storyboards []struct {
		Url              string `json:"url"`
		TemplateUrl      string `json:"templateUrl"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Count            int    `json:"count"`
		Interval         int    `json:"interval"`
		StoryboardWidth  int    `json:"storyboardWidth"`
		StoryboardHeight int    `json:"storyboardHeight"`
		StoryboardCount  int    `json:"storyboardCount"`
	} `json:"storyboards"`
	Description      string   `json:"description"`
	DescriptionHtml  string   `json:"descriptionHtml"`
	Published        int64    `json:"published"`
	PublishedText    string   `json:"publishedText"`
	Keywords         []any    `json:"keywords"`
	ViewCount        int64    `json:"viewCount"`
	LikeCount        int64    `json:"likeCount"`
	DislikeCount     int64    `json:"dislikeCount"`
	Paid             bool     `json:"paid"`
	Premium          bool     `json:"premium"`
	IsFamilyFriendly bool     `json:"isFamilyFriendly"`
	AllowedRegions   []string `json:"allowedRegions"`
	Genre            string   `json:"genre"`
	GenreUrl         string   `json:"genreUrl"`
	Author           string   `json:"author"`
	AuthorId         string   `json:"authorId"`
	AuthorUrl        string   `json:"authorUrl"`
	AuthorThumbnails []struct {
		Url    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"authorThumbnails"`
	SubCountText    string  `json:"subCountText"`
	LengthSeconds   int64   `json:"lengthSeconds"`
	AllowRatings    bool    `json:"allowRatings"`
	Rating          float32 `json:"rating"`
	IsListed        bool    `json:"isListed"`
	LiveNow         bool    `json:"liveNow"`
	IsUpcoming      bool    `json:"isUpcoming"`
	DashUrl         string  `json:"dashUrl"`
	AdaptiveFormats []struct {
		Init            string `json:"init"`
		Index           string `json:"index"`
		Bitrate         string `json:"bitrate"`
		Url             string `json:"url"`
		Itag            string `json:"itag"`
		Type            string `json:"type"`
		Clen            string `json:"clen"`
		Lmt             string `json:"lmt"`
		ProjectionType  string `json:"projectionType"`
		Fps             int    `json:"fps,omitempty"`
		Container       string `json:"container,omitempty"`
		Encoding        string `json:"encoding,omitempty"`
		AudioQuality    string `json:"audioQuality,omitempty"`
		AudioSampleRate int    `json:"audioSampleRate,omitempty"`
		AudioChannels   int    `json:"audioChannels,omitempty"`
		Resolution      string `json:"resolution,omitempty"`
		QualityLabel    string `json:"qualityLabel,omitempty"`
		ColorInfo       struct {
			Primaries               string `json:"primaries"`
			TransferCharacteristics string `json:"transferCharacteristics"`
			MatrixCoefficients      string `json:"matrixCoefficients"`
		} `json:"colorInfo,omitempty"`
	} `json:"adaptiveFormats"`
	FormatStreams []struct {
		Url          string `json:"url"`
		Itag         string `json:"itag"`
		Type         string `json:"type"`
		Quality      string `json:"quality"`
		Fps          int    `json:"fps"`
		Container    string `json:"container"`
		Encoding     string `json:"encoding"`
		Resolution   string `json:"resolution"`
		QualityLabel string `json:"qualityLabel"`
		Size         string `json:"size"`
	} `json:"formatStreams"`
	Captions []struct {
		Label        string `json:"label"`
		LanguageCode string `json:"language_code"`
		Url          string `json:"url"`
	} `json:"captions"`
	RecommendedVideos []struct {
		VideoId         string `json:"videoId"`
		Title           string `json:"title"`
		VideoThumbnails []struct {
			Quality string `json:"quality"`
			Url     string `json:"url"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"videoThumbnails"`
		Author        string `json:"author"`
		AuthorUrl     string `json:"authorUrl"`
		AuthorId      string `json:"authorId"`
		LengthSeconds int64  `json:"lengthSeconds"`
		ViewCountText string `json:"viewCountText"`
		ViewCount     int64  `json:"viewCount"`
	} `json:"recommendedVideos"`
}

func getInvidiousVideo(api, frontendUrl, videoId string) (Video, error) {
	response, err := http.Get(api + "/api/v1/videos/" + videoId)
	if err != nil {
		return Video{}, err
	}

	decodedResponse := invidiousVideoResponse{}
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	var maxResolutionVideoThumbnailUrl string
	for _, videoThumbnail := range decodedResponse.VideoThumbnails {
		if videoThumbnail.Quality == "maxres" || videoThumbnail.Quality == "maxresdefault" {
			maxResolutionVideoThumbnailUrl = videoThumbnail.Url
			break
		}
	}

	var authorThumbnailUrl string
	for _, authorThumbnail := range decodedResponse.AuthorThumbnails {
		if authorThumbnail.Width == 48 && authorThumbnail.Height == 48 {
			authorThumbnailUrl = authorThumbnail.Url
			break
		}
	}

	var processedRecommendedVideos = []Video{}
	for _, recommended := range decodedResponse.RecommendedVideos {
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
			LengthSeconds: recommended.LengthSeconds,
			ViewCount:     recommended.ViewCount,
			ViewCountText: recommended.ViewCountText,
		})
	}

	var processedCaptions = []VideoCaption{}
	for _, caption := range decodedResponse.Captions {
		processedCaptions = append(processedCaptions, VideoCaption{
			Label:    caption.Label,
			Language: caption.LanguageCode,
			Url:      api + caption.Url,
		})
	}

	embedUrl := frontendUrl + "/embed/" + decodedResponse.VideoId

	newVideo := Video{
		Title:             decodedResponse.Title,
		Id:                decodedResponse.VideoId,
		EmbedUrl:          embedUrl,
		ThumbnailUrl:      maxResolutionVideoThumbnailUrl,
		Author:            decodedResponse.Author,
		AuthorId:          decodedResponse.AuthorId,
		AuthorAvatarUrl:   authorThumbnailUrl,
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
		Captions:          processedCaptions,
		RecommendedVideos: processedRecommendedVideos,
	}

	return newVideo, nil
}
