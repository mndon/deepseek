package deepseek_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-deepseek/deepseek/internal"
)

func NewFakeServer(filePath string) *httptest.Server {
	fs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJson, _ := testdata.ReadFile(filePath)
		w.Write(respJson)
	}))
	internal.BASE_URL = fs.URL
	return fs
}
