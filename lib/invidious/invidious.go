package invidious

import (
	"encoding/json"
	"net/http"
)

func GetApiInstances() ([]string, error) {
	response, err := http.Get("https://api.invidious.io/instances.json?pretty=1&sort_by=type,users")
	if err != nil {
		return nil, err
	}

	knownInstances := [][]any{}
	json.NewDecoder(response.Body).Decode(&knownInstances)

	knownApiInstances := []string{}
	for _, instance := range knownInstances {
		mapInstance := instance[1].(map[string]any)

		if mapInstance["api"] != nil && mapInstance["api"].(bool) {
			knownApiInstances = append(knownApiInstances, mapInstance["uri"].(string))
		}
	}

	return knownApiInstances, nil
}

func ValidateInstance(instance string) (bool, error) {
	response, err := http.Get(instance + "/api/v1/stats")
	if err != nil {
		return false, err
	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		unwantedDecode := map[string]any{}
		err := json.NewDecoder(response.Body).Decode(&unwantedDecode)
		// this error is to test if we get a json body
		if err != nil {
			return false, nil
		}

		return true, nil
	}

	return false, nil
}

type GetPopularResponse []struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	VideoId         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		Url     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	LengthSeconds int    `json:"lengthSeconds"`
	ViewCount     int    `json:"viewCount"`
	Author        string `json:"author"`
	AuthorId      string `json:"authorId"`
	AuthorUrl     string `json:"authorUrl"`
	Published     int    `json:"published"`
	PublishedText string `json:"publishedText"`
}

func GetPopular(instance string) (GetPopularResponse, error) {
	// add region string (ISO 3166 country code, default is US)

	response, err := http.Get(instance + "/api/v1/popular")
	if err != nil {
		return nil, err
	}

	decodedResponse := GetPopularResponse{}
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	return decodedResponse, nil
}

type GetTrendingResponse []struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	VideoId         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		Url     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	LengthSeconds   int    `json:"lengthSeconds"`
	ViewCount       int    `json:"viewCount"`
	Author          string `json:"author"`
	AuthorId        string `json:"authorId"`
	AuthorUrl       string `json:"authorUrl"`
	Published       int    `json:"published"`
	PublishedText   string `json:"publishedText"`
	Description     string `json:"description"`
	DescriptionHtml string `json:"descriptionHtml"`
	LiveNow         bool   `json:"liveNow"`
	Paid            bool   `json:"paid"`
	Premium         bool   `json:"premium"`
}

func GetTrending(instance string) (GetTrendingResponse, error) {
	// add region string (ISO 3166 country code, default is US)

	response, err := http.Get(instance + "/api/v1/trending")
	if err != nil {
		return nil, err
	}

	decodedResponse := GetTrendingResponse{}
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	return decodedResponse, nil
}

type GetVideoResponse struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	VideoId         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		Url     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	Storyboards []struct {
		Url              string `json:"url"`
		TemplateUrl      string `json:"templateUrl"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Count            int    `json:"count"`
		Interval         int    `json:"interval"`
		StoryboardWidth  int    `json:"storyboardWidth"`
		StoryboardHeight int    `json:"storyboardHeight"`
		StoryboardCount  int    `json:"storyboardCount"`
	} `json:"storyboards"`
	Description      string   `json:"description"`
	DescriptionHtml  string   `json:"descriptionHtml"`
	Published        int      `json:"published"`
	PublishedText    string   `json:"publishedText"`
	Keywords         []any    `json:"keywords"`
	ViewCount        int      `json:"viewCount"`
	LikeCount        int      `json:"likeCount"`
	DislikeCount     int      `json:"dislikeCount"`
	Paid             bool     `json:"paid"`
	Premium          bool     `json:"premium"`
	IsFamilyFriendly bool     `json:"isFamilyFriendly"`
	AllowedRegions   []string `json:"allowedRegions"`
	Genre            string   `json:"genre"`
	GenreUrl         string   `json:"genreUrl"`
	Author           string   `json:"author"`
	AuthorId         string   `json:"authorId"`
	AuthorUrl        string   `json:"authorUrl"`
	AuthorThumbnails []struct {
		Url    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"authorThumbnails"`
	SubCountText    string  `json:"subCountText"`
	LengthSeconds   int     `json:"lengthSeconds"`
	AllowRatings    bool    `json:"allowRatings"`
	Rating          float32 `json:"rating"`
	IsListed        bool    `json:"isListed"`
	LiveNow         bool    `json:"liveNow"`
	IsUpcoming      bool    `json:"isUpcoming"`
	DashUrl         string  `json:"dashUrl"`
	AdaptiveFormats []struct {
		Init            string `json:"init"`
		Index           string `json:"index"`
		Bitrate         string `json:"bitrate"`
		Url             string `json:"url"`
		Itag            string `json:"itag"`
		Type            string `json:"type"`
		Clen            string `json:"clen"`
		Lmt             string `json:"lmt"`
		ProjectionType  string `json:"projectionType"`
		Fps             int    `json:"fps,omitempty"`
		Container       string `json:"container,omitempty"`
		Encoding        string `json:"encoding,omitempty"`
		AudioQuality    string `json:"audioQuality,omitempty"`
		AudioSampleRate int    `json:"audioSampleRate,omitempty"`
		AudioChannels   int    `json:"audioChannels,omitempty"`
		Resolution      string `json:"resolution,omitempty"`
		QualityLabel    string `json:"qualityLabel,omitempty"`
		ColorInfo       struct {
			Primaries               string `json:"primaries"`
			TransferCharacteristics string `json:"transferCharacteristics"`
			MatrixCoefficients      string `json:"matrixCoefficients"`
		} `json:"colorInfo,omitempty"`
	} `json:"adaptiveFormats"`
	FormatStreams []struct {
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
	} `json:"formatStreams"`
	Captions []struct {
		Label        string `json:"label"`
		LanguageCode string `json:"language_code"`
		Url          string `json:"url"`
	} `json:"captions"`
	RecommendedVideos []struct {
		VideoId         string `json:"videoId"`
		Title           string `json:"title"`
		VideoThumbnails []struct {
			Quality string `json:"quality"`
			Url     string `json:"url"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"videoThumbnails"`
		Author        string `json:"author"`
		AuthorUrl     string `json:"authorUrl"`
		AuthorId      string `json:"authorId"`
		LengthSeconds int    `json:"lengthSeconds"`
		ViewCountText string `json:"viewCountText"`
		ViewCount     int    `json:"viewCount"`
	} `json:"recommendedVideos"`
}

func GetVideo(instance, videoId string) (GetVideoResponse, error) {
	response, err := http.Get(instance + "/api/v1/videos/" + videoId)
	if err != nil {
		return GetVideoResponse{}, err
	}

	decodedResponse := GetVideoResponse{}
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	return decodedResponse, nil
}

// func Login(instance, email, password string) (bool, error) {
// 	data := url.Values{}
// 	data.Set("email", email)
// 	data.Set("password", password)
// 	data.Set("action", "signin")

// 	request, err := http.NewRequest("POST", instance+"/login?type=invidious", strings.NewReader(data.Encode()))
// 	if err != nil {
// 		return false, err
// 	}
// 	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	client := http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return false, err
// 	}

// 	if response.StatusCode >= 200 && response.StatusCode <= 299 {
// 		fmt.Println(request.Cookies())
// 		fmt.Println(request.Header.Get("Cookie"))
// 		fmt.Println(request.Header.Get("Set-Cookie"))
// 		fmt.Println(response.Cookies())
// 		fmt.Println(response.Header.Get("Cookie"))
// 		fmt.Println(response.Header.Get("Set-Cookie"))

// 		return true, nil
// 	}

// 	return false, nil
// }
