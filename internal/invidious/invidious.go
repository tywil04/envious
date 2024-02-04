package invidious

import (
	"fmt"
	"net/url"
)

func GetInstances() ([]Instance, error) {
	uri := "https://api.invidious.io/instances.json?pretty=1&sort_by=health,type,users,api"
	var rawInstances [][]any
	if err := JSONHTTPRequest(uri, "GET", nil, nil, &rawInstances); err != nil {
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
	if err := JSONHTTPRequest(uri, "GET", nil, nil, &video); err != nil {
		return Video{}, err
	}
	video = fixCaptions(instance, video)[0]
	video = fixHtmlDescription(instance, video)[0]
	video = proxyThumbnails(instance, video)[0]

	return video, nil
}

func GetTrendingVideos(instance Instance, options ...TrendingOption) ([]Video, error) {
	values := url.Values{}
	if len(options) > 0 {
		option := options[0]

		if option.Type != "" {
			values.Set("type", option.Type)
		}

		if option.Region != "" {
			values.Set("type", option.Region)
		}
	}

	uri := fmt.Sprintf("%s/api/v1/trending?%s", instance.ApiUrl, values.Encode())
	fmt.Println(uri)

	videos := []Video{}
	if err := JSONHTTPRequest(uri, "GET", nil, nil, &videos); err != nil {
		return []Video{}, err
	}
	videos = fixCaptions(instance, videos...)
	videos = fixHtmlDescription(instance, videos...)
	videos = proxyThumbnails(instance, videos...)

	return videos, nil
}

// sortBy: "relevance", "rating", "upload_date", "view_count"
// date: "hour", "today", "week", "month", "year"
// duration: "short", "long", "medium"
// type: (default all) "video", "playlist", "channel", "movie", "show", "all"
// features: (array containing one or more) "hd", "subtitles", "creative_commons", "3d", "live", "purchased", "4k", "360", "location", "hdr", "vr180"
// region (default US) ISO 3166 country code
func Search(instance Instance, query string, options ...SearchOption) ([]SearchItem, error) {
	values := url.Values{}
	if len(options) >= 1 {
		option := options[0]

		if option.Page > 0 {
			values.Set("page", fmt.Sprintf("%d", option.Page))
		}

		if option.SortBy != "" {
			values.Set("sort_by", option.SortBy)
		}

		if option.Date != "" {
			values.Set("date", option.Date)
		}

		if option.Duration != "" {
			values.Set("duration", option.Duration)
		}

		if option.Type != "" {
			values.Set("type", option.Type)
		}

		if option.Region != "" {
			values.Set("region", option.Region)
		}

		for _, feature := range option.Features {
			values.Add("features", feature)
		}
	}

	values.Set("q", query)

	uri := fmt.Sprintf("%s/api/v1/search?%s", instance.ApiUrl, values.Encode())
	searchItems := []SearchItem{}
	if err := JSONHTTPRequest(uri, "GET", nil, nil, &searchItems); err != nil {
		return nil, err
	}

	for index := range searchItems {
		switch searchItems[index].Type {
		case "playlist":
			searchItems[index].Videos = fixCaptions(instance, searchItems[index].Videos...)
		case "video":
			searchItems[index].Videos = fixHtmlDescription(instance, searchItems[index].Videos...)
			searchItems[index].Videos = proxyThumbnails(instance, searchItems[index].Videos...)
		case "channel":
			searchItems[index].Videos = proxyThumbnails(instance, searchItems[index].Videos...)
		}
	}

	return searchItems, nil
}
