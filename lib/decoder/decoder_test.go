package decoder_test

import (
	"testing"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/encoder"
)

func TestNullValueDecoder(t *testing.T) {
	var v = encoder.Encode(nil)
	var nullValue = decoder.Decode(v)
	if nullValue != nil {
		t.Errorf("Expected null value to be nil, but got %v", nullValue)
	}
}
