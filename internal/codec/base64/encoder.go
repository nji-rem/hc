package base64

import (
	"strings"
)

var (
	// Base64 encoding uses a segment size of 6, because 6^2 = 64 characters. It fits perfectly, and is therefore
	// the most efficient segment size.
	segmentSize = 6

	// asciiOffset starts at '@'. 0x40 in decimal is 64.
	asciiOffset = 0x40

	// maskSegmentSize gives us the last 6 bits, ensuring that we're creating groups of 6 (segmentSize).
	maskSegmentSize = 0x3f
)

func Encode(input int) string {
	var sb strings.Builder
	for i := 0; i < 2; i++ {
		// We start with the most significant bits - that's convention. Take n bits where n = segmentSize for each
		// loop, and write it to the string builder.
		//
		// An example to make everything a bit more clear:
		// "@@" = 0,
		// "@A" = 1
		// Let's use 1, or @A, as an example.
		//
		// A = 0x41 in hex, and 0b1000001 in binary. To store it in numOfBytes, we'll have to divide it by six. The
		// segment size is always the same, so we'll use segments of six.
		//
		// (2 - 1 - 0 = 1) = 1 * 6 = 6 - so a shift of 6. With other words, we'll take 0b000001.
		shift := (1 - i) * segmentSize

		// We only want to keep the data that we're going to use for this item, so we'll have to mask it by the segment
		// size (0x3f = 0b00111111). Masking it by the segment size ensures that we always get the right amount of data.
		shifted := (input >> shift) & maskSegmentSize

		// Increment the value of shifted with asciiOffset. An encoding could theoretically work without this offset,
		// but we want human-readable characters. asciiOffset (0x40 = 64) starts at @ and the characters that follow are
		// all human-readable characters (A-Z etc).
		shifted += asciiOffset

		// Append the shifted byte to the string builder.
		sb.WriteByte(byte(shifted))
	}

	return sb.String()
}
