package packet

import (
	"bytes"
	"sync"
)

var readerPool = sync.Pool{New: func() any {
	return new(Reader)
}}

var bufferPool = sync.Pool{New: func() any {
	return new(bytes.Buffer)
}}

func AcquireReader(body []byte) *Reader {
	reader := readerPool.Get().(*Reader)
	buffer := bufferPool.Get().(*bytes.Buffer)

	buffer.Reset()
	buffer.Write(body)

	reader.Buffer = buffer

	return reader
}

func ReleaseReader(reader *Reader) {
	bufferPool.Put(reader.Buffer)
	readerPool.Put(reader)
}
