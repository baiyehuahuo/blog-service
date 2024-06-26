package util

import (
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	value := "Hello"
	target := "8b1a9953c4611296a827abf8c47804d7"
	if v := EncodeMD5(value); v != target {
		t.Fatalf("EncodeMD5(%s) => %s, expected %s", value, v, target)
	}
}
