package mongorpc_test

import (
	"testing"

	"github.com/mongorpc/mongorpc"
)

func TestNullValueDecoder(t *testing.T) {
	var v = mongorpc.Encode(nil)
	var nullValue = mongorpc.Decode(v)
	if nullValue != nil {
		t.Errorf("Expected null value to be nil, but got %v", nullValue)
	}
}
