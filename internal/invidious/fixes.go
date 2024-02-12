package invidious

import (
	"regexp"
	"strings"

	"github.com/emvi/iso-639-1"

	"github.com/tywil04/tubed/internal/proxy"
)

var htmlDescriptionRegex = regexp.MustCompile(`(?m)(?:\r\n|\r|\n)`)

const captionsHmacKey = "aea7611e118cbca3a6cb0f4f86a6ad795a69040c"

func fixCaptions(instanceUrl string, videos ...Video) []Video {
	// captions dont have absolute urls, so this fixes that
	for i := range videos {
		for j := range videos[i].Captions {
			videos[i].Captions[j].Url = proxy.PathPrefix + instanceUrl + videos[i].Captions[j].Url + "&hmac_key=" + captionsHmacKey
			videos[i].Captions[j].Type = "text/vtt; charset=UTF-8"

			if videos[i].Captions[j].LanguageCode == "" {
				language := strings.TrimSuffix(videos[i].Captions[j].Label, " (auto-generated)")
				videos[i].Captions[j].LanguageCode = iso6391.CodeForName(language)
			}
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

			switch videos[i].VideoThumbnails[j].Quality {
			case "medium":
				go proxy.Preload(path)
				videos[i].TubedVideoThumbnailUrl = path
			}
		}

		for j := range videos[i].AuthorThumbnails {
			path := proxy.PathPrefix + videos[i].AuthorThumbnails[j].Url

			videos[i].AuthorThumbnails[j].Url = path

			switch videos[i].AuthorThumbnails[j].Width {
			case 48:
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
