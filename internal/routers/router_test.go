package routers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTag(t *testing.T) {
	routePaths := [][]string{
		[]string{http.MethodPost, "http://127.0.0.1:8080/api/v1/tags"},
		[]string{http.MethodDelete, "http://127.0.0.1:8080/api/v1/tags/1"},
		[]string{http.MethodPut, "http://127.0.0.1:8080/api/v1/tags/1"},
		[]string{http.MethodGet, "http://127.0.0.1:8080/api/v1/tags"},
	}
	r := NewRouter()
	for _, route := range routePaths {
		req, err := http.NewRequest(route[0], route[1], nil)
		if err != nil {
			t.Fatal(err)
		}
		record := httptest.NewRecorder()
		r.ServeHTTP(record, req)

		var (
			res = record.Result()
			buf = make([]byte, 1024)
			n   int
		)
		if n, err = res.Body.Read(buf); n != 0 || err != io.EOF {
			t.Fatalf("get message : %s, error: %v", string(buf[:n]), err)
		}

		t.Logf("tag test method %s, path %s success", route[0], route[1])
	}
}

func TestArticle(t *testing.T) {
	routePaths := [][]string{
		[]string{http.MethodPost, "http://127.0.0.1:8080/api/v1/articles"},
		[]string{http.MethodDelete, "http://127.0.0.1:8080/api/v1/articles/1"},
		[]string{http.MethodPut, "http://127.0.0.1:8080/api/v1/articles/1"},
		[]string{http.MethodGet, "http://127.0.0.1:8080/api/v1/articles/1"},
		[]string{http.MethodGet, "http://127.0.0.1:8080/api/v1/articles"},
	}
	r := NewRouter()
	for _, route := range routePaths {
		req, err := http.NewRequest(route[0], route[1], nil)
		if err != nil {
			t.Fatal(err)
		}
		record := httptest.NewRecorder()
		r.ServeHTTP(record, req)

		var (
			res = record.Result()
			buf = make([]byte, 1024)
			n   int
		)
		if n, err = res.Body.Read(buf); n != 0 || err != io.EOF {
			t.Fatalf("get message : %s, error: %v", string(buf[:n]), err)
		}

		t.Logf("article test method %s, path %s success", route[0], route[1])
	}
}
