package base64_test

import (
	"hc/pkg/encoding/base64"
	"testing"
)

var testCases = map[int]string{
	0:       "@@",
	1:       "@A",
	266:     "DJ",
	14:      "@N",
	666:     "JZ",
	4234324: "qT",
}

func TestEncode(t *testing.T) {
	for n, expected := range testCases {
		encodedValue := base64.Encode(n)
		if encodedValue != expected {
			t.Errorf("expected base64 value %s for integer %d, got %s instead", expected, n, encodedValue)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = base64.Encode(1024)
	}
}
