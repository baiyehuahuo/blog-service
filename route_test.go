package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := getRouter()
	record := httptest.NewRecorder()
	r.ServeHTTP(record, req)

	var (
		res  = record.Result()
		buf  = make([]byte, 1024)
		n    int
		body gin.H
	)
	if n, err = res.Body.Read(buf); err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(buf[:n], &body); err != nil {
		t.Fatal(err)
	}

	pongMessage := gin.H{"message": "pong"}
	if !reflect.DeepEqual(body, pongMessage) {
		t.Fatal("got", body, "expected", pongMessage)
	}

	t.Log("pong test success")
}
