package errcode

import "testing"

func TestNewError(t *testing.T) {
	err := NewError(100, "handling")
	if err.code != 100 || err.msg != "handling" {
		t.Fatalf("err.code = %d, err.msg = %s", err.code, err.msg)
	}

	err = ServerError
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("insert repeat success err.code = %d, err.msg = %s", err.code, err.msg)
		}
	}()
	_ = NewError(err.code, err.msg)
}
