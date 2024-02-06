package invidious

import (
	"regexp"

	"github.com/tywil04/tubed/internal/proxy"
)

var htmlDescriptionRegex = regexp.MustCompile(`(?m)(?:\r\n|\r|\n)`)

func fixCaptions(instanceUrl string, videos ...Video) []Video {
	// captions dont have absolute urls, so this fixes that
	for i := range videos {
		for j := range videos[i].Captions {
			videos[i].Captions[j].Url = instanceUrl + videos[i].Captions[j].Url
		}

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = fixCaptions(instanceUrl, videos[i].RecommendedVideos...)
		}
	}

	return videos
}

func fixHtmlDescription(instanceUrl string, videos ...Video) []Video {
	// html descriptions dont use breaks for newlines
	for i := range videos {
		videos[i].DescriptionHtml = htmlDescriptionRegex.ReplaceAllString(videos[i].DescriptionHtml, "<br/>")

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = fixHtmlDescription(instanceUrl, videos[i].RecommendedVideos...)
		}
	}
	return videos
}

func proxyThumbnails(videos ...Video) []Video {
	// set urls for images so they proxy through go
	for i := range videos {
		for j := range videos[i].VideoThumbnails {
			path := proxy.PathPrefix + videos[i].VideoThumbnails[j].Url

			videos[i].VideoThumbnails[j].Url = path

			if videos[i].VideoThumbnails[j].Quality == "medium" {
				go proxy.Preload(path)
				videos[i].TubedVideoThumbnailUrl = path
			}
		}

		for j := range videos[i].AuthorThumbnails {
			path := proxy.PathPrefix + videos[i].AuthorThumbnails[j].Url

			videos[i].AuthorThumbnails[j].Url = path

			if videos[i].AuthorThumbnails[j].Width == 48 {
				go proxy.Preload(path)
				videos[i].TubedAuthorThumbnailUrl = path
			}
		}

		if len(videos[i].RecommendedVideos) > 0 {
			videos[i].RecommendedVideos = proxyThumbnails(videos[i].RecommendedVideos...)
		}
	}
	return videos
}
