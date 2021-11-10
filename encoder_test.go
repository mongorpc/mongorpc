package mongorpc_test

import (
	"testing"

	"github.com/mongorpc/mongorpc"
)

func TestStringValueEncoder(t *testing.T) {
	e := mongorpc.Encoder{}
	v := "Hello World!"
	r := e.Encode(v)
	if r.GetStringValue() != v {
		t.Errorf("Expected %s, got %s", v, r.GetStringValue())
	}
}

func TestIntValueEncoder(t *testing.T) {
	e := mongorpc.Encoder{}
	v := 12345
	r := e.Encode(v)
	if int(r.GetIntegerValue()) != v {
		t.Errorf("Expected %d, got %d", v, r.GetIntegerValue())
	}
}

func TestBoolValueEncoder(t *testing.T) {
	e := mongorpc.Encoder{}
	v := true
	r := e.Encode(v)
	if r.GetBoolValue() != v {
		t.Errorf("Expected %t, got %t", v, r.GetBoolValue())
	}
}
