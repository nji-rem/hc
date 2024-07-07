package base64

const (
	// SegmentSize Base64 encoding uses a segment size of 6, because 6^2 = 64 characters. It fits perfectly, and is therefore
	// the most efficient segment size.
	SegmentSize = 6

	// AsciiOffset starts at '@'. 0x40 in decimal is 64.
	AsciiOffset = 0x40

	// SegmentSizeMask gives us the last 6 bits, ensuring that we're creating groups of 6 (SegmentSize).
	SegmentSizeMask = 0x3f
)
