package shared

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// this is currently unused
//
// func jsonEncodeBody(body any) (io.Reader, error) {
// 	var rawBody *bytes.Buffer
// 	err := json.NewEncoder(rawBody).Encode(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rawBody, nil
// }

// Make a http GET request.
// url is the url to GET.
// body is an io.Reader and can be generated with functions like jsonEncodeBody.
// headers is a map, the key is Header and the value is the value to set.
// value is a pointer to a value that will be filled with the response.
func HttpGetJson(url string, body io.Reader, headers map[string]string, value any) error {
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&value)
	if err != nil {
		return err
	}

	return nil
}

// Make a http GET request and return the resulting body as a string.
// url is the url to GET.
// body is an io.Reader and can be generated with functions like jsonEncodeBody.
// headers is a map, the key is Header and the value is the value to set.
func HttpGetString(url string, body io.Reader, headers map[string]string) (string, error) {
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	stringBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(stringBody), nil
}

// Make a http GET request and capture and return the redirect location.
// url is the url to GET.
// body is an io.Reader and can be generated with functions like jsonEncodeBody.
// headers is a map, the key is Header and the value is the value to set.
func HttpGetRedirect(url string, body io.Reader, headers map[string]string) (string, error) {
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return "", nil
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := http.Client{}

	var redirectUrl string
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		redirect, err := request.Response.Location()
		redirectUrl = redirect.String()
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return "", nil
	}

	return redirectUrl, nil
}
