// Package that interfaces with Piped and Invidious to get YouTube videos
package providers

import (
	"errors"

	"github.com/tywil04/tubed/internal/providers/invidious"
	"github.com/tywil04/tubed/internal/providers/piped"
	"github.com/tywil04/tubed/internal/providers/shared"
)

// Get list of instances from provider.
// provider is either 'invidious' or 'piped'.
func GetInstancesApi(provider string) ([]map[string]string, error) {
	switch provider {
	case "invidious":
		return invidious.GetInstances()
	case "piped":
		return piped.GetInstances()
	default:
		return []map[string]string{}, errors.New("no instance provider selected")
	}
}

// Get a providers frontend url from its api url.
// provider is either 'invidious' or 'piped'.
func GetInstanceFrontend(provider string, api string) (string, error) {
	switch provider {
	case "invidious":
		return invidious.GetInstanceFrontend(api)
	case "piped":
		return piped.GetInstanceFrontend(api)
	default:
		return "", errors.New("no instance provider selected")
	}
}

// Gets the trending videos from provider.
// provider is either 'invidious' or 'piped'.
func GetTrending(provider, api, region string) ([]shared.Video, error) {
	switch provider {
	case "invidious":
		return invidious.GetTrending(api, region)
	case "piped":
		return piped.GetTrending(api, region)
	default:
		return []shared.Video{}, errors.New("no instance provider selected")
	}
}

// Gets the video from provider via videoId.
// provider is either 'invidious' or 'piped'.
// videoId is a YouTube video id.
func GetVideo(provider, api, frontend, videoId string) (shared.Video, error) {
	switch provider {
	case "invidious":
		return invidious.GetVideo(api, frontend, videoId)
	case "piped":
		return piped.GetVideo(api, frontend, videoId)
	default:
		return shared.Video{}, errors.New("no instance provider selected")
	}
}
