package invidious

import (
	"fmt"
	"net/url"
)

type TrendingOption struct {
	Type   string `json:"type"`
	Region string `json:"region"`
}

func (s *Session) GetTrendingVideos(option ...TrendingOption) ([]Video, error) {
	values := url.Values{}
	if len(option) > 0 {
		if option[0].Type != "" {
			values.Set("type", option[0].Type)
		}

		if option[0].Region != "" {
			values.Set("type", option[0].Region)
		}
	}

	endpoint := fmt.Sprintf("/api/v1/trending?%s", values.Encode())

	videos := []Video{}
	if err := s.makeRequest(endpoint, "GET", nil, nil, &videos); err != nil {
		return []Video{}, err
	}
	videos = fixCaptions(s.instance.ApiUrl, videos...)
	videos = fixHtmlDescription(s.instance.ApiUrl, videos...)
	videos = proxyThumbnails(videos...)

	return videos, nil
}
