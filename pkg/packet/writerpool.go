package packet

import (
	"github.com/valyala/bytebufferpool"
	"sync"
)

var writerPool = sync.Pool{New: func() any {
	return new(Writer)
}}

func AcquireWriter() *Writer {
	writer := writerPool.Get().(*Writer)
	writer.Buffer = bytebufferpool.Get()

	return writer
}

func ReleaseWriter(writer *Writer) {
	bytebufferpool.Put(writer.Buffer)

	writer.Buffer = nil
	writerPool.Put(writer)
}
