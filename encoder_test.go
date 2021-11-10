package mongorpc_test

import (
	"testing"

	"github.com/mongorpc/mongorpc"
)

func TestStringValueEncoder(t *testing.T) {
	v := "Hello World!"
	r := mongorpc.Encode(v)
	if r.GetStringValue() != v {
		t.Errorf("Expected %s, got %s", v, r.GetStringValue())
	}
}

func TestIntValueEncoder(t *testing.T) {
	v := 12345
	r := mongorpc.Encode(v)
	if int(r.GetIntegerValue()) != v {
		t.Errorf("Expected %d, got %d", v, r.GetIntegerValue())
	}
}

func TestBoolValueEncoder(t *testing.T) {
	v := true
	r := mongorpc.Encode(v)
	if r.GetBoolValue() != v {
		t.Errorf("Expected %t, got %t", v, r.GetBoolValue())
	}
}
