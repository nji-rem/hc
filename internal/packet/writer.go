package packet

import (
	"github.com/valyala/bytebufferpool"
	"hc/pkg/encoding/base64"
	"hc/pkg/encoding/vl64"
	"io"
)

type Writer struct {
	Buffer *bytebufferpool.ByteBuffer
}

func (w *Writer) AppendHeader(header int) error {
	encoded := base64.Encode(header)
	if _, err := w.Buffer.WriteString(encoded); err != nil {
		return err
	}

	return nil
}

func (w *Writer) AppendInt(integer int) error {
	if _, err := w.Buffer.Write(vl64.Encode(integer)); err != nil {
		return err
	}

	return nil
}

func (w *Writer) AppendString(str string) error {
	if _, err := w.Buffer.WriteString(str); err != nil {
		return err
	}

	// Each str ends with chr(2)
	if err := w.Buffer.WriteByte(2); err != nil {
		return err
	}

	return nil
}

func (w *Writer) WriteTo(writer io.Writer) (n int64, err error) {
	if err := w.Buffer.WriteByte(1); err != nil {
		return 0, err
	}

	defer func() {
		offset := len(w.Buffer.B) - 1
		w.Buffer.B = w.Buffer.B[:offset]
	}()

	bytesWritten, err := writer.Write(w.Buffer.B)
	if err != nil {
		return 0, err
	}

	return int64(bytesWritten), nil
}
