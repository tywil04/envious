package assetProxy

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

var cachePath = os.TempDir() + "/" + "TubedCache"

func hashKey(key string) string {
	h := md5.New()
	io.WriteString(h, key)
	return hex.EncodeToString(h.Sum(nil))
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
