package memo

type Memo struct{ requests chan request }

type request struct {
	key      string
	response chan result
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.Monitor(f)
	return memo
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	result := <-response
	return result.value, result.err
}

func (m *Memo) Monitor(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests {
		_entry := cache[req.key]
		if _entry == nil {
			_entry = &entry{ready: make(chan struct{})}
			cache[req.key] = _entry
			go _entry.call(f, req.key)
		}
		go _entry.ship(req.response)
	}
}

func (m *Memo) Close() { close(m.requests) }

func (e *entry) ship(response chan<- result) {
	<-e.ready
	response <- e.res

}
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}
