package connection

import (
	"errors"
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/session"
	"io"
)

var ErrInstance = errors.New("handler relies on gnet.Conn as io.Writer")
var ErrCtxInstance = errors.New("expected an instance of connection.Context")

type CreateSessionOnNewConnection struct {
	SessionStore session.Store
	Pool         *session.Pool
}

func (s CreateSessionOnNewConnection) Handle(w io.Writer) error {
	conn, ok := w.(gnet.Conn)
	if !ok {
		return ErrInstance
	}

	// not a memory leak! there's a shutdown handler that removes session data. it needs an integration test
	// to test this clean up functionality.
	sess := s.Pool.Acquire()

	if err := s.SessionStore.Add(sess); err != nil {
		return err
	}

	conn.SetContext(connection.NewContext(sess.ID))

	log.Info().Msgf("Created session with id %s", sess.ID)

	return nil
}

type DeleteSessionOnConnectionDestroyed struct {
	SessionStore session.Store
	Pool         *session.Pool
}

func (d DeleteSessionOnConnectionDestroyed) Handle(w io.Writer) error {
	conn, ok := w.(gnet.Conn)
	if !ok {
		return ErrInstance
	}

	ctx, ok := conn.Context().(connection.Context)
	if !ok {
		return ErrCtxInstance
	}

	sessionObj, err := d.SessionStore.Delete(ctx.SessionID())
	if err != nil {
		return err
	}

	log.Info().Msgf("Removed session with id %s", ctx.SessionID())

	d.Pool.Release(sessionObj)

	return nil
}
