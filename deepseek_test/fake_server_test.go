package deepseek_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/mndon/deepseek/internal"
)

func NewFakeServer(filePath string) *httptest.Server {
	fs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJson, _ := testdata.ReadFile(filePath)
		w.Write(respJson)
	}))
	internal.BASE_URL = fs.URL
	return fs
}

func NewFakeSteamServer(filePath string) *httptest.Server {
	respJson, _ := testdata.ReadFile(filePath)
	next := dataIter(respJson)

	fs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			msg, found := next()
			if !found {
				break
			}

			w.Write([]byte(msg))
			w.(http.Flusher).Flush()

			if msg == "data: [DONE]" {
				break
			}
		}
	}))

	internal.BASE_URL = fs.URL
	return fs
}

func dataIter(content []byte) func() (string, bool) {
	data := strings.Split(string(content), "data: ")
	i := 0
	length := len(data)
	return func() (string, bool) {
		for {
			if i >= length {
				return "", false
			}
			if data[i] == "" {
				i++
				continue
			} else {
				res := data[i]
				i++
				return fmt.Sprintf("data: %s", res), true
			}
		}
	}
}
