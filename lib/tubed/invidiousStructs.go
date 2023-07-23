package tubed

type invidiousVideoThumbnail struct {
	Quality string `json:"quality"`
	Url     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type invidiousStoryboard struct {
	Url              string `json:"url"`
	TemplateUrl      string `json:"templateUrl"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	Count            int    `json:"count"`
	Interval         int    `json:"interval"`
	StoryboardWidth  int    `json:"storyboardWidth"`
	StoryboardHeight int    `json:"storyboardHeight"`
	StoryboardCount  int    `json:"storyboardCount"`
}

type invidiousAuthorThumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type invidiousAdaptiveFormatsColorInfo struct {
	Primaries               string `json:"primaries"`
	TransferCharacteristics string `json:"transferCharacteristics"`
	MatrixCoefficients      string `json:"matrixCoefficients"`
}

type invidiousAdaptiveFormat struct {
	Init            string                            `json:"init"`
	Index           string                            `json:"index"`
	Bitrate         string                            `json:"bitrate"`
	Url             string                            `json:"url"`
	Itag            string                            `json:"itag"`
	Type            string                            `json:"type"`
	Clen            string                            `json:"clen"`
	Lmt             string                            `json:"lmt"`
	ProjectionType  string                            `json:"projectionType"`
	Fps             int                               `json:"fps,omitempty"`
	Container       string                            `json:"container,omitempty"`
	Encoding        string                            `json:"encoding,omitempty"`
	AudioQuality    string                            `json:"audioQuality,omitempty"`
	AudioSampleRate int                               `json:"audioSampleRate,omitempty"`
	AudioChannels   int                               `json:"audioChannels,omitempty"`
	Resolution      string                            `json:"resolution,omitempty"`
	QualityLabel    string                            `json:"qualityLabel,omitempty"`
	ColorInfo       invidiousAdaptiveFormatsColorInfo `json:"colorInfo,omitempty"`
}

type invidiousFormatStream struct {
	Url          string `json:"url"`
	Itag         string `json:"itag"`
	Type         string `json:"type"`
	Quality      string `json:"quality"`
	Fps          int    `json:"fps"`
	Container    string `json:"container"`
	Encoding     string `json:"encoding"`
	Resolution   string `json:"resolution"`
	QualityLabel string `json:"qualityLabel"`
	Size         string `json:"size"`
}

type invidiousCaption struct {
	Label        string `json:"label"`
	LanguageCode string `json:"language_code"`
	Url          string `json:"url"`
}

type invidiousRecommendedVideo struct {
	VideoId         string                    `json:"videoId"`
	Title           string                    `json:"title"`
	VideoThumbnails []invidiousVideoThumbnail `json:"videoThumbnails"`
	Author          string                    `json:"author"`
	AuthorUrl       string                    `json:"authorUrl"`
	AuthorId        string                    `json:"authorId"`
	LengthSeconds   int64                     `json:"lengthSeconds"`
	ViewCountText   string                    `json:"viewCountText"`
	ViewCount       int64                     `json:"viewCount"`
}

type invidiousTrendingResponse []struct {
	Type            string                    `json:"type"`
	Title           string                    `json:"title"`
	VideoId         string                    `json:"videoId"`
	VideoThumbnails []invidiousVideoThumbnail `json:"videoThumbnails"`
	LengthSeconds   int64                     `json:"lengthSeconds"`
	ViewCount       int64                     `json:"viewCount"`
	Author          string                    `json:"author"`
	AuthorId        string                    `json:"authorId"`
	AuthorUrl       string                    `json:"authorUrl"`
	Published       int64                     `json:"published"`
	PublishedText   string                    `json:"publishedText"`
	Description     string                    `json:"description"`
	DescriptionHtml string                    `json:"descriptionHtml"`
	LiveNow         bool                      `json:"liveNow"`
	Paid            bool                      `json:"paid"`
	Premium         bool                      `json:"premium"`
}

type invidiousVideoResponse struct {
	Type              string                      `json:"type"`
	Title             string                      `json:"title"`
	VideoId           string                      `json:"videoId"`
	VideoThumbnails   []invidiousVideoThumbnail   `json:"videoThumbnails"`
	Storyboards       []invidiousStoryboard       `json:"storyboards"`
	Description       string                      `json:"description"`
	DescriptionHtml   string                      `json:"descriptionHtml"`
	Published         int64                       `json:"published"`
	PublishedText     string                      `json:"publishedText"`
	Keywords          []any                       `json:"keywords"`
	ViewCount         int64                       `json:"viewCount"`
	LikeCount         int64                       `json:"likeCount"`
	DislikeCount      int64                       `json:"dislikeCount"`
	Paid              bool                        `json:"paid"`
	Premium           bool                        `json:"premium"`
	IsFamilyFriendly  bool                        `json:"isFamilyFriendly"`
	AllowedRegions    []string                    `json:"allowedRegions"`
	Genre             string                      `json:"genre"`
	GenreUrl          string                      `json:"genreUrl"`
	Author            string                      `json:"author"`
	AuthorId          string                      `json:"authorId"`
	AuthorUrl         string                      `json:"authorUrl"`
	AuthorThumbnails  []invidiousAuthorThumbnail  `json:"authorThumbnails"`
	SubCountText      string                      `json:"subCountText"`
	LengthSeconds     int64                       `json:"lengthSeconds"`
	AllowRatings      bool                        `json:"allowRatings"`
	Rating            float32                     `json:"rating"`
	IsListed          bool                        `json:"isListed"`
	LiveNow           bool                        `json:"liveNow"`
	IsUpcoming        bool                        `json:"isUpcoming"`
	DashUrl           string                      `json:"dashUrl"`
	AdaptiveFormats   []invidiousAdaptiveFormat   `json:"adaptiveFormats"`
	FormatStreams     []invidiousFormatStream     `json:"formatStreams"`
	Captions          []invidiousCaption          `json:"captions"`
	RecommendedVideos []invidiousRecommendedVideo `json:"recommendedVideos"`
}
