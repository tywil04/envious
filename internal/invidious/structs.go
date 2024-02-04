package invidious

type videoThumbnail struct {
	Quality string `json:"quality"`
	Url     string `json:"url"`
	Width   int32  `json:"width"`
	Height  int32  `json:"height"`
}

type authorThumbnail struct {
	Url    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

type adaptiveFormat struct {
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

type formatStream struct {
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

type caption struct {
	Label        string `json:"label"`
	LanguageCode string `json:"languageCode"`
	Url          string `json:"url"`
}

type Video struct {
	Title             string            `json:"title"`
	VideoId           string            `json:"videoId"`
	VideoThumbnails   []videoThumbnail  `json:"videoThumbnails"`
	Description       string            `json:"description"`
	DescriptionHtml   string            `json:"descriptionHtml"`
	Published         int64             `json:"published"`
	PublishedText     string            `json:"publishedText"`
	Keywords          []string          `json:"keywords"`
	ViewCount         int64             `json:"viewCount"`
	LikeCount         int32             `json:"likeCount"`
	DislikeCount      int32             `json:"dislikeCount"`
	Paid              bool              `json:"paid"`
	Premium           bool              `json:"premium"`
	IsFamilyFriendly  bool              `json:"isFamilyFriendly"`
	AllowedRegions    []string          `json:"allowedRegions"`
	Genre             string            `json:"genre"`
	GenreUrl          string            `json:"genreUrl"`
	Author            string            `json:"author"`
	AuthorId          string            `json:"authorId"`
	AuthorUrl         string            `json:"authorUrl"`
	AuthorThumbnails  []authorThumbnail `json:"authorThumbnails"`
	SubCountText      string            `json:"subCountText"`
	LengthSeconds     int32             `json:"lengthSeconds"`
	AllowRatings      bool              `json:"allowRatings"`
	Rating            float32           `json:"rating"`
	IsListed          bool              `json:"isListed"`
	LiveNow           bool              `json:"liveNow"`
	IsUpcoming        bool              `json:"isUpcoming"`
	PremiereTimestamp int64             `json:"premiereTimestamp"`
	DashUrl           string            `json:"dashUrl"`
	AdaptiveFormats   []adaptiveFormat  `json:"adaptiveFormats"`
	FormatStreams     []formatStream    `json:"formatStreams"`
	Captions          []caption         `json:"captions"`
	RecommendedVideos []Video           `json:"recommendedVideos"`

	// these fields are added by tubed so the frontend can do less work
	// fields are added while running fixes
	TubedVideoThumbnailUrl  string `json:"tubedVideoThumbnailUrl"`
	TubedAuthorThumbnailUrl string `json:"tubedAuthorThumbnailUrl"`
}

type SearchOption struct {
	Page     int32    `json:"page"`
	SortBy   string   `json:"sortBy"`
	Date     string   `json:"date"`
	Duration string   `json:"duration"`
	Type     string   `json:"type"`
	Features []string `json:"features"`
	Region   string   `json:"region"`
}

type SearchItem struct {
	// shared
	Type      string `json:"type"`
	Title     string `json:"title"`
	VideoId   string `json:"videoId"`
	Author    string `json:"author"`
	AuthorId  string `json:"authorId"`
	AuthorUrl string `json:"authorUrl"`
	ViewCount int64  `json:"viewCount"`

	// video
	VideoThumbnails []videoThumbnail `json:"videoThumbnails"`
	Description     string           `json:"description"`
	DescriptionHtml string           `json:"descriptionHtml"`
	Published       int64            `json:"published"`
	PublishedText   string           `json:"publishedText"`
	LengthSeconds   int32            `json:"lengthSeconds"`
	LiveNow         bool             `json:"liveNow"`
	Paid            bool             `json:"paid"`
	Premium         bool             `json:"premium"`

	// playlist
	AuthorVerified    bool    `json:"authorVerified"`
	PlaylistId        string  `json:"playlistId"`
	PlaylistThumbnail string  `json:"playlistThumbnail"`
	Videos            []Video `json:"videos"`

	// channel
	AuthorThumbnails []authorThumbnail `json:"authorThumbnails"`
	AutoGenerated    bool              `json:"autoGenerated"`
	SubCount         int32             `json:"subCount"`
	VideoCount       int32             `json:"videoCount"`
}

type Instance struct {
	ApiUrl string `json:"apiUrl"`
	Cors   bool   `json:"cors"`
	Region string `json:"region"`
}

type TrendingOption struct {
	Type   string `json:"type"`
	Region string `json:"region"`
}
