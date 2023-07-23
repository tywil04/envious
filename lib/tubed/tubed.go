// Package that interfaces with Piped and Invidious to get YouTube videos
package tubed

import (
	"errors"
)

// Get list of instances from provider.
// provider is either 'invidious' or 'piped'.
func GetInstancesApi(provider string) ([]map[string]string, error) {
	switch provider {
	case "invidious":
		return getInvidiousInstances()
	case "piped":
		return getPipedInstances()
	default:
		return nil, errors.New("no instance provider selected")
	}
}

// Get a providers frontend url from its api url.
// provider is either 'invidious' or 'piped'.
func GetInstanceFrontend(provider string, api string) (string, error) {
	switch provider {
	case "invidious":
		return getInvidiousInstanceFrontend(api)
	case "piped":
		return getPipedInstanceFrontend(api)
	default:
		return "", errors.New("no instance provider selected")
	}
}

// Gets the trending videos from provider.
// provider is either 'invidious' or 'piped'.
func GetTrending(provider, api, region string) ([]Video, error) {
	switch provider {
	case "invidious":
		return getInvidiousTrending(api, region)
	case "piped":
		return getPipedTrending(api, region)
	default:
		return nil, errors.New("no instance provider selected")
	}
}

// Gets the video from provider via videoId.
// provider is either 'invidious' or 'piped'.
// videoId is a YouTube video id.
func GetVideo(provider, api, frontend, videoId string) (Video, error) {
	switch provider {
	case "invidious":
		return getInvidiousVideo(api, frontend, videoId)
	case "piped":
		return getPipedVideo(api, frontend, videoId)
	default:
		return Video{}, errors.New("no instance provider selected")
	}
}
