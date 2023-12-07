package invidious

import (
	"fmt"
	"net/url"
)

func fixCaptions(instance Instance, videos ...Video) []Video {
	// captions dont have absolute urls, so this fixes that
	for videoIndex, video := range videos {
		for captionIndex := range video.Captions {
			videos[videoIndex].Captions[captionIndex].Url = instance.ApiUrl + videos[videoIndex].Captions[captionIndex].Url
		}
	}
	return videos
}

func GetInstances() ([]Instance, error) {
	uri := "https://api.invidious.io/instances.json?pretty=1&sort_by=health,type,users,api"
	var rawInstances [][]any
	if err := JSONHTTPRequest(uri, "GET", &rawInstances); err != nil {
		return nil, err
	}

	instances := []Instance{}
	for _, unprocessedInstance := range rawInstances {
		data, ok := unprocessedInstance[1].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("failed to process instances data")
		}

		if data["api"] != nil && data["type"] == "https" {
			instances = append(instances, Instance{
				ApiUrl: data["uri"].(string),
				Cors:   data["cors"].(bool),
				Region: data["region"].(string),
			})
		}
	}

	return instances, nil
}

func GetVideo(instance Instance, videoId string) (Video, error) {
	uri := fmt.Sprintf("%s/api/v1/videos/%s", instance.ApiUrl, videoId)
	video := Video{}
	if err := JSONHTTPRequest(uri, "GET", &video); err != nil {
		return Video{}, err
	}
	video = fixCaptions(instance, video)[0]
	return video, nil
}

func GetTrendingVideos(instance Instance) ([]Video, error) {
	uri := fmt.Sprintf("%s/api/v1/trending", instance.ApiUrl)
	videos := []Video{}
	if err := JSONHTTPRequest(uri, "GET", &videos); err != nil {
		return []Video{}, err
	}
	videos = fixCaptions(instance, videos...)
	return videos, nil
}

func GetPopularVideos(instance Instance) ([]Video, error) {
	uri := fmt.Sprintf("%s/api/v1/popular", instance.ApiUrl)
	videos := []Video{}
	if err := JSONHTTPRequest(uri, "GET", &videos); err != nil {
		return []Video{}, err
	}
	videos = fixCaptions(instance, videos...)
	return videos, nil
}

// sortBy: "relevance", "rating", "upload_date", "view_count"
// date: "hour", "today", "week", "month", "year"
// duration: "short", "long", "medium"
// type: (default all) "video", "playlist", "channel", "movie", "show", "all"
// features: (array containing one or more) "hd", "subtitles", "creative_commons", "3d", "live", "purchased", "4k", "360", "location", "hdr", "vr180"
// region (default US) ISO 3166 country code
func Search(instance Instance, query string, options ...SearchOptions) ([]SearchItem, error) {
	values := url.Values{}
	values.Set("q", query)
	if len(options) >= 1 {
		option := options[0]
		values.Set("page", fmt.Sprintf("%d", option.Page))
		values.Set("sort_by", option.SortBy)
		values.Set("date", option.Date)
		values.Set("duration", option.Duration)
		values.Set("type", option.Type)
		values.Set("region", option.Region)
		for _, feature := range option.Features {
			values.Add("features", feature)
		}
	}

	uri := fmt.Sprintf("%s/api/v1/search/%s?%s", instance.ApiUrl, query, values.Encode())
	searchItems := []SearchItem{}
	if err := JSONHTTPRequest(uri, "GET", &searchItems); err != nil {
		return nil, err
	}

	for _, searchItem := range searchItems {
		if searchItem.Type == "playlist" {
			searchItem.Videos = fixCaptions(instance, searchItem.Videos...)
		}
	}

	return searchItems, nil
}
