package db

import (
	"io/fs"
	"os"
)

func createDirIfNotExist(dirPath string, filePerm fs.FileMode) error {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dirPath, 0777)
		} else {
			return err
		}
	}
	return nil
}

func createFileIfNotExist(filePath, defaultContents string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			configFile, err := os.Create(filePath)
			if err != nil {
				return err
			}
			configFile.Write([]byte(defaultContents))
			configFile.Close()
		} else {
			return err
		}
	}
	return nil
}
