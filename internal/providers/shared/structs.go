package shared

type Caption struct {
	Label    string `json:"label"`
	Language string `json:"language"`
	Url      string `json:"url"`
}

type Video struct {
	Title             string    `json:"title"`
	Id                string    `json:"id"`
	EmbedUrl          string    `json:"embedUrl"`
	ThumbnailUrl      string    `json:"thumbnailUrl"`
	Author            string    `json:"author"`
	AuthorId          string    `json:"authorId"`
	AuthorAvatarUrl   string    `json:"authorAvatarUrl"`
	ShortDescription  string    `json:"shortDescription"`
	Description       string    `json:"description"`
	Published         int64     `json:"published"`
	PublishedText     string    `json:"publishedText"`
	Genre             string    `json:"genre"`
	SubCountText      string    `json:"subCountText"`
	LengthSeconds     int64     `json:"lengthSeconds"`
	AllowRatings      bool      `json:"allowRatings"`
	Rating            float32   `json:"rating"`
	ViewCount         int64     `json:"viewCount"`
	ViewCountText     string    `json:"viewCountText"`
	LikeCount         int64     `json:"likeCount"`
	DislikeCount      int64     `json:"dislikeCount"`
	IsListed          bool      `json:"isListed"`
	IsUpcoming        bool      `json:"isUpcoming"`
	IsShort           bool      `json:"isShort"`
	IsPaid            bool      `json:"isPaid"`
	IsPremium         bool      `json:"isPremium"`
	IsFamilyFriendly  bool      `json:"isFamilyFriendly"`
	IsLiveNow         bool      `json:"liveNow"`
	Captions          []Caption `json:"captions"`
	RecommendedVideos []Video   `json:"recommendedVideos"`
}
