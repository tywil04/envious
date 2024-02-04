package proxy

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

var cachePath = os.TempDir() + "/" + "TubedCache"

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
