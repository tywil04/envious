package invidious

type videoThumbnail struct {
	Quality string `json:"quality" toml:"quality"`
	Url     string `json:"url" toml:"url"`
	Width   int32  `json:"width" toml:"width"`
	Height  int32  `json:"height" toml:"height"`
}

type authorThumbnail struct {
	Url    string `json:"url" toml:"url"`
	Width  int32  `json:"width" toml:"width"`
	Height int32  `json:"height" toml:"height"`
}

type adaptiveFormat struct {
	Index           string `json:"index" toml:"index"`
	Bitrate         string `json:"bitrate" toml:"bitrate"`
	Init            string `json:"init" toml:"init"`
	Url             string `json:"url" toml:"url"`
	ITag            string `json:"iTag" toml:"iTag"`
	Type            string `json:"type" toml:"type"`
	Clen            string `json:"clen" toml:"clen"`
	Lmt             string `json:"lmt" toml:"lmt"`
	ProjectionType  string `json:"projectionType" toml:"projectionType"`
	Fps             int32  `json:"fps" toml:"fps"`
	Container       string `json:"container" toml:"container"`
	Encoding        string `json:"encoding" toml:"encoding"`
	QualityLabel    string `json:"qualityLabel" toml:"qualityLabel"`
	Resolution      string `json:"resolution" toml:"resolution"`
	AudioQuality    string `json:"audioQuality" toml:"audioQuality"`
	AudioSampleRate int32  `json:"audioSampleRate" toml:"audioSampleRate"`
	AudioChannels   int32  `json:"audioChannels" toml:"audioChannels"`
}

type formatStream struct {
	Url          string `json:"url" toml:"url"`
	ITag         string `json:"iTag" toml:"iTag"`
	Type         string `json:"type" toml:"type"`
	Quality      string `json:"quality" toml:"quality"`
	Fps          int32  `json:"fps" toml:"fps"`
	Container    string `json:"container" toml:"container"`
	Encoding     string `json:"encoding" toml:"encoding"`
	QualityLabel string `json:"qualityLabel" toml:"qualityLabel"`
	Resultion    string `json:"resultion" toml:"resultion"`
	Size         string `json:"size" toml:"size"`
}

type caption struct {
	Label        string `json:"label" toml:"label"`
	LanguageCode string `json:"languageCode" toml:"languageCode"`
	Url          string `json:"url" toml:"url"`
}

type Video struct {
	Title             string            `json:"title" toml:"title"`
	VideoId           string            `json:"videoId" toml:"videoId"`
	VideoThumbnails   []videoThumbnail  `json:"videoThumbnails" toml:"videoThumbnails"`
	Description       string            `json:"description" toml:"description"`
	DescriptionHtml   string            `json:"descriptionHtml" toml:"descriptionHtml"`
	Published         int64             `json:"published" toml:"published"`
	PublishedText     string            `json:"publishedText" toml:"publishedText"`
	Keywords          []string          `json:"keywords" toml:"keywords"`
	ViewCount         int64             `json:"viewCount" toml:"viewCount"`
	LikeCount         int32             `json:"likeCount" toml:"likeCount"`
	DislikeCount      int32             `json:"dislikeCount" toml:"dislikeCount"`
	Paid              bool              `json:"paid" toml:"paid"`
	Premium           bool              `json:"premium" toml:"premium"`
	IsFamilyFriendly  bool              `json:"isFamilyFriendly" toml:"isFamilyFriendly"`
	AllowedRegions    []string          `json:"allowedRegions" toml:"allowedRegions"`
	Genre             string            `json:"genre" toml:"genre"`
	GenreUrl          string            `json:"genreUrl" toml:"genreUrl"`
	Author            string            `json:"author" toml:"author"`
	AuthorId          string            `json:"authorId" toml:"authorId"`
	AuthorUrl         string            `json:"authorUrl" toml:"authorUrl"`
	AuthorThumbnails  []authorThumbnail `json:"authorThumbnails" toml:"authorThumbnails"`
	SubCountText      string            `json:"subCountText" toml:"subCountText"`
	LengthSeconds     int32             `json:"lengthSeconds" toml:"lengthSeconds"`
	AllowRatings      bool              `json:"allowRatings" toml:"allowRatings"`
	Rating            float32           `json:"rating" toml:"rating"`
	IsListed          bool              `json:"isListed" toml:"isListed"`
	LiveNow           bool              `json:"liveNow" toml:"liveNow"`
	IsUpcoming        bool              `json:"isUpcoming" toml:"isUpcoming"`
	PremiereTimestamp int64             `json:"premiereTimestamp" toml:"premiereTimestamp"`
	DashUrl           string            `json:"dashUrl" toml:"dashUrl"`
	HlsUrl            string            `json:"hlsUrl" toml:"hlsUrl"`
	AdaptiveFormats   []adaptiveFormat  `json:"adaptiveFormats" toml:"adaptiveFormats"`
	FormatStreams     []formatStream    `json:"formatStreams" toml:"formatStreams"`
	Captions          []caption         `json:"captions" toml:"captions"`
	RecommendedVideos []Video           `json:"recommendedVideos" toml:"recommendedVideos"`
}

type SearchOptions struct {
	Page     int32    `json:"page" toml:"page"`
	SortBy   string   `json:"sortBy" toml:"sortBy"`
	Date     string   `json:"date" toml:"date"`
	Duration string   `json:"duration" toml:"duration"`
	Type     string   `json:"type" toml:"type"`
	Features []string `json:"features" toml:"features"`
	Region   string   `json:"region" toml:"region"`
}

type SearchItem struct {
	// shared
	Type      string `json:"type" toml:"type"`
	Title     string `json:"title" toml:"title"`
	VideoId   string `json:"videoId" toml:"videoId"`
	Author    string `json:"author" toml:"author"`
	AuthorId  string `json:"authorId" toml:"authorId"`
	AuthorUrl string `json:"authorUrl" toml:"authorUrl"`
	ViewCount int64  `json:"viewCount" toml:"viewCount"`

	// video
	VideoThumbnails []videoThumbnail `json:"videoThumbnails" toml:"videoThumbnails"`
	Description     string           `json:"description" toml:"description"`
	DescriptionHtml string           `json:"descriptionHtml" toml:"descriptionHtml"`
	Published       int64            `json:"published" toml:"published"`
	PublishedText   string           `json:"publishedText" toml:"publishedText"`
	LengthSeconds   int32            `json:"lengthSeconds" toml:"lengthSeconds"`
	LiveNow         bool             `json:"liveNow" toml:"liveNow"`
	Paid            bool             `json:"paid" toml:"paid"`
	Premium         bool             `json:"premium" toml:"premium"`

	// playlist
	AuthorVerified    bool    `json:"authorVerified" toml:"authorVerified"`
	PlaylistId        string  `json:"playlistId" toml:"playlistId"`
	PlaylistThumbnail string  `json:"playlistThumbnail" toml:"playlistThumbnail"`
	Videos            []Video `json:"videos" toml:"videos"`

	// channel `json:"//" toml:"//"` only
	AuthorThumbnails []authorThumbnail `json:"authorThumbnails" toml:"authorThumbnails"`
	AutoGenerated    bool              `json:"autoGenerated" toml:"autoGenerated"`
	SubCount         int32             `json:"subCount" toml:"subCount"`
	VideoCount       int32             `json:"videoCount" toml:"videoCount"`
}

type Instance struct {
	ApiUrl string `json:"apiUrl" toml:"apiUrl"`
	Cors   bool   `json:"cors" toml:"cors"`
	Region string `json:"region" toml:"region"`
}
