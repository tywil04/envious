// Tubed: basic package to allow for config
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Provider         string `json:"provider"`
	InstanceApi      string `json:"instanceApi"`
	InstanceFrontend string `json:"instanceFrontend"`
	Region           string `json:"region"`
	Configured       bool   `json:"configured"`
}

const (
	ConfigDirectory = "/Tubed"
	ConfigFilePath  = "/config.json"
)

var Stored Config

var empty = Config{
	Provider:         "piped",
	InstanceApi:      "",
	InstanceFrontend: "",
	Region:           "US",
	Configured:       false,
}

func createConfigIfNotExist() error {
	userConfigDirectory, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configDirectory := userConfigDirectory + ConfigDirectory
	configFilePath := configDirectory + ConfigFilePath

	_, err = os.Stat(configDirectory)
	if os.IsNotExist(err) {
		os.Mkdir(configDirectory, 0777)
	} else if err != nil {
		return err
	}

	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		fmt.Println(err)

		configFile, err := os.Create(configFilePath)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(configFile).Encode(empty); err != nil {
			return err
		}

		configFile.Close()
	} else if err != nil {
		return err
	}

	return nil
}

func Load() error {
	userConfigDirectory, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	if err := createConfigIfNotExist(); err != nil {
		return err
	}

	configFile, err := os.Open(userConfigDirectory + ConfigDirectory + ConfigFilePath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	if err := json.NewDecoder(configFile).Decode(&Stored); err != nil {
		return err
	}

	return nil
}

func Offload() error {
	userConfigDirectory, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configFile, err := os.OpenFile(userConfigDirectory+ConfigDirectory+ConfigFilePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer configFile.Close()

	return json.NewEncoder(configFile).Encode(Stored)
}
