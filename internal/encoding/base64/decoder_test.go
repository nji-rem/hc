package base64_test

import (
	"hc/internal/encoding/base64"
	"testing"
)

var testCasesDecode = map[string]int{
	"@C": 3,
	"@@": 0,
	"DU": 277,
	"KU": 725,
	"ZF": 1670,
	"JZ": 666,
	"gm": 2541,
}

func TestDecode(t *testing.T) {
	for n, expected := range testCasesDecode {
		decodedValue := base64.Decode([]byte(n))
		if decodedValue != expected {
			t.Errorf("Expected decodedValue to be expected %d, got %d instead", expected, decodedValue)
		}
	}
}
