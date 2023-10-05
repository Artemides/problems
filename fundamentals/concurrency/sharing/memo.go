package sharing

import (
	"sync"
)

type Memo struct {
	f     Func
	mutex sync.Mutex
	memo  map[string]*entry
}

type Func func(key string) (interface{}, error)
type entry struct {
	res   result
	ready chan struct{}
}
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, memo: make(map[string]*entry)}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mutex.Lock()
	_entry := m.memo[key]
	if _entry != nil {
		m.mutex.Unlock()
		<-_entry.ready
		return _entry.res.value, _entry.res.err
	}

	_entry = &entry{ready: make(chan struct{})}
	m.mutex.Unlock()
	_entry.res.value, _entry.res.err = m.f(key)
	close(_entry.ready)
	return _entry.res.value, _entry.res.err
}
