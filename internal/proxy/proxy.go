package proxy

import (
	"io"
	"net/http"
	"strings"
)

const PathPrefix = "/assetproxy/"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, PathPrefix) {
			proxyUrl := strings.TrimPrefix(r.URL.Path, PathPrefix)

			data, ok := GetFromCache(proxyUrl)
			if !ok {
				response, err := http.Get(proxyUrl)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer response.Body.Close()

				w.Header().Set("Content-Type", response.Header.Get("Content-Type"))

				data, err := io.ReadAll(response.Body)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				_, err = w.Write(data)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				WriteToCache(proxyUrl, data)
			} else {
				_, err := w.Write(data)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
