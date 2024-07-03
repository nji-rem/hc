package base64

func Decode(data []byte) (out int) {
	for i := 0; i < len(data); i++ {
		// Remove the ASCII offset. The Habbo client adds an offset to the ASCII table (starting at "@" = 0x40), but
		// we have to remove that here due to the fact that we're interested in the original byte values.
		byteVal := data[i] - AsciiOffset

		// Move every existing bit to the left. This is necessary, because we got to make room for byteVal.
		out = out << 6

		byteVal = byteVal & SegmentSizeMask

		// Add byteVal to our integer
		out = out | int(byteVal)
	}

	return
}
