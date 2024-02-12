package invidious

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const instancesApiUrl = "https://api.invidious.io/instances.json?pretty=1&sort_by=health,type,users,api"

type Instance struct {
	ApiUrl string `json:"apiUrl"`
	Region string `json:"region"`
}

func GetInstances() ([]*Instance, error) {
	response, err := http.Get(instancesApiUrl)
	if err != nil {
		return nil, err
	}

	var rawInstances [][]any
	err = json.NewDecoder(response.Body).Decode(&rawInstances)
	if err != nil {
		return nil, err
	}

	var instances []*Instance
	for _, unprocessedInstance := range rawInstances {
		data, ok := unprocessedInstance[1].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("failed to process instances data")
		}

		if data["api"] != nil && data["type"] == "https" {
			instances = append(instances, &Instance{
				ApiUrl: data["uri"].(string),
				Region: data["region"].(string),
			})
		}
	}

	return instances, nil
}
