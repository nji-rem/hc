package session

import (
	"fmt"
	apiSession "hc/api/session"
	"sync"
)

type Store struct {
	sessions sync.Map
}

func (s *Store) Add(session *apiSession.Bag) error {
	s.sessions.Store(session.ID, session)

	return nil
}

func (s *Store) Get(id string) (*apiSession.Bag, error) {
	item, ok := s.sessions.Load(id)
	if !ok {
		return nil, fmt.Errorf("session with id %s does not exist", id)
	}

	return item.(*apiSession.Bag), nil
}

func (s *Store) Delete(id string) (*apiSession.Bag, error) {
	item, ok := s.sessions.Load(id)
	if !ok {
		return nil, fmt.Errorf("session with id %s does not exist", id)
	}

	s.sessions.Delete(id)

	return item.(*apiSession.Bag), nil
}
