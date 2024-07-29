package request

import "sync"

// Body contains the viewmodel body.
type Body struct {
	raw        []byte
	parsedBody any

	rw sync.RWMutex
}

func (b *Body) Raw() []byte {
	return b.raw
}

func (b *Body) SetRaw(raw []byte) {
	b.rw.Lock()
	b.raw = raw
	b.rw.Unlock()
}

func (b *Body) Parsed() any {
	b.rw.RLock()
	defer b.rw.RUnlock()

	if b.parsedBody == nil {
		return nil
	}

	return b.parsedBody
}

func (b *Body) SetParsedBody(item any) {
	b.rw.Lock()
	b.parsedBody = item
	b.rw.Unlock()
}
