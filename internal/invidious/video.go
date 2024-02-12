package invidious

import "fmt"

type videoVideoThumbnail struct {
	Quality string `json:"quality"`
	Url     string `json:"url"`
	Width   int32  `json:"width"`
	Height  int32  `json:"height"`
}

type videoAuthorThumbnail struct {
	Url    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

type videoAdaptiveFormat struct {
	Index           string `json:"index"`
	Bitrate         string `json:"bitrate"`
	Init            string `json:"init"`
	Url             string `json:"url"`
	ITag            string `json:"iTag"`
	Type            string `json:"type"`
	Clen            string `json:"clen"`
	Lmt             string `json:"lmt"`
	ProjectionType  string `json:"projectionType"`
	Fps             int32  `json:"fps"`
	Container       string `json:"container"`
	Encoding        string `json:"encoding"`
	QualityLabel    string `json:"qualityLabel"`
	Resolution      string `json:"resolution"`
	AudioQuality    string `json:"audioQuality"`
	AudioSampleRate int32  `json:"audioSampleRate"`
	AudioChannels   int32  `json:"audioChannels"`
}

type videoFormatStream struct {
	Url          string `json:"url"`
	ITag         string `json:"iTag"`
	Type         string `json:"type"`
	Quality      string `json:"quality"`
	Fps          int32  `json:"fps"`
	Container    string `json:"container"`
	Encoding     string `json:"encoding"`
	QualityLabel string `json:"qualityLabel"`
	Resultion    string `json:"resultion"`
	Size         string `json:"size"`
}

type videoCaption struct {
	Label        string `json:"label"`
	LanguageCode string `json:"languageCode"`
	Url          string `json:"url"`
	Type         string `json:"type"`
}

type Video struct {
	Title             string                 `json:"title"`
	VideoId           string                 `json:"videoId"`
	VideoThumbnails   []videoVideoThumbnail  `json:"videoThumbnails"`
	Description       string                 `json:"description"`
	DescriptionHtml   string                 `json:"descriptionHtml"`
	Published         int64                  `json:"published"`
	PublishedText     string                 `json:"publishedText"`
	Keywords          []string               `json:"keywords"`
	ViewCount         int64                  `json:"viewCount"`
	LikeCount         int32                  `json:"likeCount"`
	DislikeCount      int32                  `json:"dislikeCount"`
	Paid              bool                   `json:"paid"`
	Premium           bool                   `json:"premium"`
	IsFamilyFriendly  bool                   `json:"isFamilyFriendly"`
	AllowedRegions    []string               `json:"allowedRegions"`
	Genre             string                 `json:"genre"`
	GenreUrl          string                 `json:"genreUrl"`
	Author            string                 `json:"author"`
	AuthorId          string                 `json:"authorId"`
	AuthorUrl         string                 `json:"authorUrl"`
	AuthorThumbnails  []videoAuthorThumbnail `json:"authorThumbnails"`
	SubCountText      string                 `json:"subCountText"`
	LengthSeconds     int32                  `json:"lengthSeconds"`
	AllowRatings      bool                   `json:"allowRatings"`
	Rating            float32                `json:"rating"`
	IsListed          bool                   `json:"isListed"`
	LiveNow           bool                   `json:"liveNow"`
	IsUpcoming        bool                   `json:"isUpcoming"`
	PremiereTimestamp int64                  `json:"premiereTimestamp"`
	DashUrl           string                 `json:"dashUrl"`
	AdaptiveFormats   []videoAdaptiveFormat  `json:"adaptiveFormats"`
	FormatStreams     []videoFormatStream    `json:"formatStreams"`
	Captions          []videoCaption         `json:"captions"`
	RecommendedVideos []Video                `json:"recommendedVideos"`

	// these fields are added by tubed so the frontend can do less work
	// fields are added while running fixes
	TubedVideoThumbnailUrl  string `json:"tubedVideoThumbnailUrl"`
	TubedAuthorThumbnailUrl string `json:"tubedAuthorThumbnailUrl"`
}

func (s *Session) GetVideo(videoId string) (Video, error) {
	endpoint := fmt.Sprintf("/api/v1/videos/%s", videoId)

	video := Video{}
	err := s.makeRequest(endpoint, "GET", nil, nil, &video)
	if err != nil {
		return Video{}, err
	}
	video = fixCaptions(s.instance.ApiUrl, video)[0]
	video = fixHtmlDescription(s.instance.ApiUrl, video)[0]
	video = proxyThumbnails(video)[0]

	return video, nil
}
