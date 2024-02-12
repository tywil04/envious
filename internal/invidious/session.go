package invidious

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; rv:122.0) Gecko/20100101 Firefox/122.0"

type Session struct {
	instance *Instance
}

func NewSession(instance *Instance) *Session {
	return &Session{
		instance: instance,
	}
}

// body should be a stucture that can be marshalled by encoding/json
// responseBody should be a structure that can be marshalled by encoding/json
func (s *Session) makeRequest(endpoint, method string, headers map[string]string, body any, responseBody any) error {
	bodyBuffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(bodyBuffer).Encode(&body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, s.instance.ApiUrl+endpoint, bodyBuffer)
	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Control-Allow-Origin", "*")
	request.Header.Add("User-Agent", userAgent)
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return err
	}

	return nil
}
