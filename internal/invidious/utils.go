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

func JSONHTTPRequest(uri, method string, value any, options ...Options) error {
	var body io.Reader
	var headers = map[string]string{}
	if len(options) >= 1 {
		body = options[0].Body
		headers = options[0].Headers
	}

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
