package tubed

import (
	"errors"
)

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
