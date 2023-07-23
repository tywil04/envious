package tubed

type pipedAudioStream struct {
	Bitrate    int    `json:"bitrate"`
	Codec      string `json:"codec"`
	Format     string `json:"format"`
	IndexEnd   int    `json:"indexEnd"`
	IndexStart int    `json:"indexStart"`
	InitStart  int    `json:"initStart"`
	InitEnd    int    `json:"initEnd"`
	MimeType   string `json:"mimeType"`
	Quality    string `json:"quality"`
	Url        string `json:"url"`
	VideoOnly  bool   `json:"videoOnly"`
}

type pipedSubtitle struct {
	AutoGenerated bool   `json:"autoGenerated"`
	Code          string `json:"code"`
	MimeType      string `json:"mimeType"`
	Name          string `json:"name"`
	Url           string `json:"url"`
}

type pipedRelatedStream struct {
	Duration         int64  `json:"duration"`
	Thumbnail        string `json:"thumbnail"`
	Title            string `json:"title"`
	UploadedDate     string `json:"uploadedDate"`
	UploaderName     string `json:"uploaderName"`
	UploaderAvatar   string `json:"uploaderAvatar"`
	UploaderUrl      string `json:"uploaderUrl"`
	UploaderVerified bool   `json:"uploaderVerified"`
	Uploaded         int64  `json:"uploaded"`
	ShortDescription string `json:"shortDescription"`
	Url              string `json:"url"`
	Views            int64  `json:"views"`
	IsShort          bool   `json:"isShort"`
}

type pipedVideoStream struct {
	Bitrate    int    `json:"bitrate"`
	Codec      string `json:"codec"`
	Format     string `json:"format"`
	Fps        int    `json:"fps"`
	Height     int    `json:"height"`
	Width      int    `json:""`
	IndexEnd   int    `json:"indexEnd"`
	IndexStart int    `json:"indexStart"`
	InitStart  int    `json:"initStart"`
	InitEnd    int    `json:"initEnd"`
	MimeType   string `json:"mimeType"`
	Quality    string `json:"quality"`
	Url        string `json:"url"`
	VideoOnly  bool   `json:"videoOnly"`
}

type pipedTrendingResponse []struct {
	Duration         int64  `json:"duration"`
	Thumbnail        string `json:"thumbnail"`
	Title            string `json:"title"`
	UploadedDate     string `json:"uploadedDate"`
	UploaderName     string `json:"uploaderName"`
	UploaderAvatar   string `json:"uploaderAvatar"`
	UploaderUrl      string `json:"uploaderUrl"`
	UploaderVerified bool   `json:"uploaderVerified"`
	Uploaded         int64  `json:"uploaded"`
	ShortDescription string `json:"shortDescription"`
	Url              string `json:"url"`
	Views            int64  `json:"views"`
	IsShort          bool   `json:"isShort"`
}

type pipedVideoResponse struct {
	AudioStreams     []pipedAudioStream   `json:"audioStreams"`
	Dash             string               `json:"dash"`
	Description      string               `json:"description"`
	Dislikes         int64                `json:"dislikes"`
	Duration         int64                `json:"duration"`
	Hls              string               `json:"hls"`
	LbryId           string               `json:"lbryId"`
	Likes            int64                `json:"Likes"`
	Livestream       bool                 `json:"livestream"`
	ProxyUrl         string               `json:"proxyUrl"`
	RelatedStreams   []pipedRelatedStream `json:"relatedStreams"`
	Subtitles        []pipedSubtitle      `json:"subtitles"`
	ThumbnailUrl     string               `json:"thumbnailUrl"`
	Title            string               `json:"title"`
	UploadedDate     string               `json:"uploadedDate"`
	Uploader         string               `json:"uploader"`
	UploaderAvatar   string               `json:"uploaderAvatar"`
	UploaderUrl      string               `json:"uploaderUrl"`
	UploaderVerified bool                 `json:"uploaderVerified"`
	Uploaded         int64                `json:"uploaded"`
	Url              string               `json:"url"`
	Views            int64                `json:"views"`
	IsShort          bool                 `json:"isShort"`
	VideoStreams     []pipedVideoStream   `json:"videoStreams"`
}