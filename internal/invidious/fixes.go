package invidious

import (
	"regexp"

	"github.com/tywil04/tubed/internal/proxy"
)

var htmlDescriptionRegex = regexp.MustCompile(`(?m)(?:\r\n|\r|\n)`)

func fixCaptions(instance Instance, videos ...Video) []Video {
	// captions dont have absolute urls, so this fixes that
	for i := range videos {
		for j := range videos[i].Captions {
			videos[i].Captions[j].Url = instance.ApiUrl + videos[i].Captions[j].Url
		}

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = fixCaptions(instance, videos[i].RecommendedVideos...)
		}
	}

	return videos
}

func fixHtmlDescription(instance Instance, videos ...Video) []Video {
	// html descriptions dont use breaks for newlines
	for i := range videos {
		videos[i].DescriptionHtml = htmlDescriptionRegex.ReplaceAllString(videos[i].DescriptionHtml, "<br/>")

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = fixHtmlDescription(instance, videos[i].RecommendedVideos...)
		}
	}
	return videos
}

func proxyThumbnails(instance Instance, videos ...Video) []Video {
	// set urls for images so they proxy through go
	for i := range videos {
		for j := range videos[i].VideoThumbnails {
			videos[i].VideoThumbnails[j].Url = proxy.PathPrefix + videos[i].VideoThumbnails[j].Url
		}

		for j := range videos[i].AuthorThumbnails {
			videos[i].AuthorThumbnails[j].Url = proxy.PathPrefix + videos[i].AuthorThumbnails[j].Url
		}

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = proxyThumbnails(instance, videos[i].RecommendedVideos...)
		}
	}
	return videos
}
