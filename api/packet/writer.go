package packet

type Writer interface {
	// AppendHeader base64 encodes the header and appends it to the buffer.
	AppendHeader(header int) error

	// AppendInt appends an integer to the buffer. The writer is responsible for VL64 encoding the value.
	AppendInt(integer int) error

	// AppendString appends the string as it is to the buffer.
	AppendString(str string) error

	Bytes() ([]byte, error)
}
