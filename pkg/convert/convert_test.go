package convert

import "testing"

func TestStrTo_Int(t *testing.T) {
	x := -10
	strX := StrTo("-10")
	if v, err := strX.Int(); err != nil || v != x {
		t.Errorf("expect %d, got %d, error: %v", x, v, err)
	}
}

func TestStrTo_MustInt(t *testing.T) {
	x := -10
	strX := StrTo("-10")
	if v := strX.MustInt(); v != x {
		t.Errorf("expect %d, got %d", x, v)
	}
	strX = StrTo("ax")
	if v := strX.MustInt(); v != 0 {
		t.Errorf("expect %d, got %d", 0, v)
	}
}

func TestStrTo_UInt32(t *testing.T) {
	x := uint32(10)
	strX := StrTo("10")
	if v, err := strX.UInt32(); err != nil || v != x {
		t.Errorf("expect %d, got %d, error: %v", x, v, err)
	}
}

func TestStrTo_MustUInt32(t *testing.T) {
	x := uint32(10)
	strX := StrTo("10")
	if v := strX.MustUInt32(); v != x {
		t.Errorf("expect %d, got %d", x, v)
	}
	strX = StrTo("ax")
	if v := strX.MustUInt32(); v != 0 {
		t.Errorf("expect %d, got %d", 0, v)
	}
}

func TestStrTo_String(t *testing.T) {
	x := "10"
	strX := StrTo(x)
	if v := strX.String(); v != x {
		t.Errorf("expect %s, got %s", x, v)
	}
}
