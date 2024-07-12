package connection

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"hc/api/connection"
	"io"
	"reflect"
)

type TrafficParser struct{}

func (t TrafficParser) Parse(reader io.Reader, request *connection.Request) error {
	conn, ok := reader.(gnet.Conn)
	if !ok {
		// I don't expect this to happen, so this reflection shouldn't happen in real life.
		readerType := reflect.TypeOf(reader)
		return fmt.Errorf("expected instance of gnet.Conn, got %s instead", readerType.String())
	}

	buffer, err := conn.Next(-1)
	if err != nil {
		return err
	}

	if len(buffer) < 5 {
		return fmt.Errorf("expected buffer size of at least 5, got length %d", len(buffer))
	}

	request.Header = string(buffer[3:5])
	request.Body = buffer[5:]

	return nil
}
