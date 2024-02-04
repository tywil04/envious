package proxy

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"

	"github.com/djherbis/atime"
)

const (
	pathSeperator = string(os.PathSeparator)

	fileBecomesOldAfter = 14 // days
)

var cachePath = os.TempDir() + pathSeperator + "TubedCache"

func init() {
	_, err := os.Stat(cachePath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(cachePath, os.ModeAppend)
		} else {
			panic(err)
		}
	}

	err = cleanupOldCacheEntries()
	if err != nil {
		panic(err)
	}
}

func cleanupOldCacheEntries() error {
	files, err := os.ReadDir(cachePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return err
		}

		lastAccessed := atime.Get(info)
		oneMonthAgo := time.Now().AddDate(0, 0, -fileBecomesOldAfter)

		if lastAccessed.Before(oneMonthAgo) {
			filePath := cachePath + pathSeperator + info.Name()
			os.Remove(filePath)
		}
	}

	return nil
}

func hashKey(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetFromCache(key string) ([]byte, bool) {
	key = hashKey(key)
	data, err := os.ReadFile(cachePath + "/" + key)
	if err != nil {
		return nil, false
	}
	return data, true
}

func WriteToCache(key string, data []byte) bool {
	key = hashKey(key)
	err := os.WriteFile(cachePath+"/"+key, data, os.ModePerm)
	return err == nil
}

func DeleteFromCache(key string) bool {
	key = hashKey(key)
	err := os.Remove(cachePath + "/" + key)
	return err == nil
}
