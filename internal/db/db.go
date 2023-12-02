// Tubed: basic package to allow for config
package db

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	config = map[string]any{}

	configDir  string
	configFile string
)

func init() {
	userConfigDir, _ := os.UserConfigDir()

	configDir = userConfigDir + "/Tubed"
	configFile = configDir + "/data.toml"
}

func Read() error {
	err := createDirIfNotExist(configDir, 0777)
	if err != nil {
		return err
	}

	err = createFileIfNotExist(configFile)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		return err
	}

	return nil
}

func Write() error {
	file, err := os.OpenFile(configFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(&config); err != nil {
		return err
	}

	return nil
}

func Set(key string, value any) {
	levels := strings.Split(key, ".")
	levelsLen := len(levels)

	last := config
	for index, level := range levels {
		if index == levelsLen-1 {
			last[level] = value
		} else {
			last[level] = map[string]any{}
			last = last[level].(map[string]any)
		}
	}
}

func Get[T any](key string) T {
	levels := strings.Split(key, ".")
	levelsLen := len(levels)

	var value T
	last := config
	for index, level := range levels {
		if index == levelsLen-1 {
			value, _ = last[level].(T)
		} else {
			var ok bool
			last, ok = last[level].(map[string]any)
			if !ok {
				// not set, break so default is returned
				break
			}
		}
	}

	return value
}

func GetWithDefault[T any](key string, defaultValue any) T {
	levels := strings.Split(key, ".")
	levelsLen := len(levels)

	var value T
	var valueOk bool
	last := config
	for index, level := range levels {
		if index == levelsLen-1 {
			value, valueOk = last[level].(T)
		} else {
			var ok bool
			last, ok = last[level].(map[string]any)
			if !ok {
				// not set, break so default is returned
				break
			}
		}
	}

	if !valueOk {
		var defaultValueOk bool
		value, defaultValueOk = defaultValue.(T)
		if !defaultValueOk {
			panic("tried to get from db with default value that doesn't match return type")
		}
	}

	return value
}
