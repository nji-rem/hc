package packet

import (
	"bytes"
	"hc/pkg/encoding/base64"
	"hc/pkg/encoding/vl64"
	"io"
)

type Reader struct {
	Buffer *bytes.Buffer
}

func (r *Reader) Short() (int, error) {
	buf := r.Buffer.Next(2)

	return base64.Decode(buf), nil
}

func (r *Reader) Int() (int, error) {
	v, length, err := vl64.Decode(r.Buffer.Bytes())
	if err != nil {
		return 0, err
	}

	r.Buffer.Next(length)

	return v, nil
}

func (r *Reader) String() (string, error) {
	if r.Buffer.Len() < 2 {
		return "", io.EOF
	}

	packetLength := base64.Decode(r.Buffer.Next(2))

	return string(r.Buffer.Next(packetLength)), nil
}
