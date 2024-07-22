package packet

import (
	"bytes"
	"hc/pkg/encoding/base64"
	"io"
)

type Reader struct {
	Buffer *bytes.Buffer
}

func (r *Reader) String() (string, error) {
	if r.Buffer.Len() < 2 {
		return "", io.EOF
	}

	packetLength := base64.Decode(r.Buffer.Next(2))

	return string(r.Buffer.Next(packetLength)), nil
}
