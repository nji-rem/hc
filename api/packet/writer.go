package packet

import "io"

type Writer interface {
	// WriterTo ensures that the byte buffer can be written to e.g. a network stream.
	io.WriterTo

	// AppendHeader base64 encodes the header and appends it to the buffer.
	AppendHeader(header int) error

	// AppendInt appends an integer to the buffer. The writer is responsible for VL64 encoding the value.
	AppendInt(integer int) error

	// AppendString appends the string as it is to the buffer.
	AppendString(str string) error
}
