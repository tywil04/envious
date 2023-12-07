package assetProxy

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

type Proxy struct{}

func NewProxy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyUrl := strings.TrimPrefix(r.URL.Path, "/")

		data, ok := GetFromCache(proxyUrl)
		if !ok {
			request, err := http.NewRequest(r.Method, proxyUrl, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			client := http.Client{}
			response, err := client.Do(request)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer response.Body.Close()

			buffer := bytes.NewBuffer([]byte{})
			_, err = io.Copy(buffer, response.Body)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			_, err = w.Write(buffer.Bytes())
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ok := WriteToCache(proxyUrl, buffer.Bytes())
			if !ok {
				w.Write([]byte("Failed to write file to cache."))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			_, err := w.Write(data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
}
