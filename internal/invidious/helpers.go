package invidious

import (
	"encoding/json"
	"io"
	"net/http"
)

type Options struct {
	Body    io.Reader
	Headers map[string]string
}

func JSONHTTPRequest(uri, method string, headers map[string]string, body io.Reader, value any) error {
	request, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Access-Control-Allow-Origin", "*")
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
