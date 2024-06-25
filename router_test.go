package main

import (
	"blog-service/internal/routers"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

// todo how to test
func TestTagCreate(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("name", "Golang")
	_ = writer.WriteField("created_by", "fwf")
	if err := writer.Close(); err != nil {
		t.Fatal("multipart.writer close err:", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/v1/tags", payload)
	if err != nil {
		t.Fatalf("Build request failed, err: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r := routers.NewRouter()
	record := httptest.NewRecorder()
	r.ServeHTTP(record, req)
	res := record.Result()
	if res.StatusCode != http.StatusOK {
		msg := make([]byte, 1024)
		n, _ := res.Body.Read(msg)
		t.Fatalf("create tag failed, status code: %v, msg: %v", res.Status, string(msg[:n]))
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
	r := routers.NewRouter()
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
