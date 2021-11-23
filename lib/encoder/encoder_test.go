package encoder_test

import (
	"testing"

	"github.com/mongorpc/mongorpc/lib/encoder"
)

func TestStringValueEncoder(t *testing.T) {
	v := "Hello World!"
	r := encoder.Encode(v)
	if r.GetStringValue() != v {
		t.Errorf("Expected %s, got %s", v, r.GetStringValue())
	}
}

func TestIntValueEncoder(t *testing.T) {
	v := int64(12345)
	r := encoder.Encode(v)
	if (r.GetInteger64Value()) != v {
		t.Errorf("Expected %d, got %d", v, r.GetInteger64Value())
	}
}

func TestBoolValueEncoder(t *testing.T) {
	v := true
	r := encoder.Encode(v)
	if r.GetBooleanValue() != v {
		t.Errorf("Expected %t, got %t", v, r.GetBooleanValue())
	}
}
