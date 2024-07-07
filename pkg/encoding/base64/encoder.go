package base64

// Encode takes an integer and encodes it to a base64 string.
func Encode(input int) string {
	msb := (input >> SegmentSize & SegmentSizeMask) + AsciiOffset
	lsb := (input & SegmentSizeMask) + AsciiOffset

	return string([]byte{byte(msb), byte(lsb)})
}
