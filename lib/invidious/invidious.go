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
	return response.StatusCode >= 200 && response.StatusCode <= 299, nil
}
