package invidious

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	VideoID         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		URL     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	LengthSeconds int    `json:"lengthSeconds"`
	ViewCount     int64  `json:"viewCount"`
	Author        string `json:"author"`
	AuthorID      string `json:"authorId"`
	AuthorURL     string `json:"authorUrl"`
	Published     int64  `json:"published"`
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
	VideoID         string `json:"videoId"`
	VideoThumbnails []struct {
		Quality string `json:"quality"`
		URL     string `json:"url"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"videoThumbnails"`
	LengthSeconds   int    `json:"lengthSeconds"`
	ViewCount       int64  `json:"viewCount"`
	Author          string `json:"author"`
	AuthorID        string `json:"authorId"`
	AuthorURL       string `json:"authorUrl"`
	Published       int64  `json:"published"`
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

func Login(instance, email, password string) (bool, error) {
	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)
	data.Set("action", "signin")

	request, err := http.NewRequest("POST", instance+"/login?type=invidious", strings.NewReader(data.Encode()))
	if err != nil {
		return false, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		fmt.Println(request.Cookies())
		fmt.Println(request.Header.Get("Cookie"))
		fmt.Println(request.Header.Get("Set-Cookie"))
		fmt.Println(response.Cookies())
		fmt.Println(response.Header.Get("Cookie"))
		fmt.Println(response.Header.Get("Set-Cookie"))

		return true, nil
	}

	return false, nil
}
