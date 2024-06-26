package upload

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"strings"
	"testing"
)

func TestGetFileExt(t *testing.T) {
	inputs := []string{
		"hello.png",
		"hello.pwd",
		"hello.jpn",
		"www",
	}
	for _, v := range inputs {
		var target string
		lastIdx := strings.LastIndex(v, ".")
		if lastIdx != -1 {
			target = v[lastIdx:]
		} else {
			target = ""
		}
		if output := GetFileExt(v); target != output {
			t.Errorf("GetFileExt failed, expected %s, got %s", target, output)
		}
	}
}

// how to check it logically ?
func TestGetFileName(t *testing.T) {
	filename := util.EncodeMD5("aaa") + ".txt"
	if filename != GetFileName("aaa.txt") {
		t.Errorf("GetFileName failed, expected %s, got %s", GetFileName("aaa.txt"), filename)
	}
}

func TestGetSavePath(t *testing.T) {
	// it must be checked by go test ./... because it needs global init
	defer func() {
		_ = recover()
	}()
	if GetSavePath() != global.AppSetting.UploadSavePath {
		t.Fatalf("GetSavePath failed")
	}
}

func TestCheckSavePath(t *testing.T) {
	if CheckSavePath("file.go") || !CheckSavePath("www.go") {
		t.Fatalf("CheckSavePath failed")
	}
}

func TestCheckContainExt(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	if !CheckContainExt(TypeImage, "www.png") {
		t.Fatalf("CheckContainExt failed")
	}
	if !CheckContainExt(TypeImage, "www.jpg") {
		t.Fatalf("CheckContainExt failed")
	}
	if !CheckContainExt(TypeImage, "www.jpeg") {
		t.Fatalf("CheckContainExt failed")
	}
	if CheckContainExt(TypeImage, "www.jpp") {
		t.Fatalf("CheckContainExt failed")
	}
}

func TestCheckPermission(t *testing.T) {
	// todo how to check it another?
	if CheckPermission("file.go") {
		t.Fatalf("CheckPermission failed")
	}
}

// todo how to check it safely?
func TestCreateSavePath(t *testing.T) {}
func TestCheckMaxSize(t *testing.T)   {}
func TestSaveFile(t *testing.T)       {}
