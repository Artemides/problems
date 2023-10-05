package memov2

type request struct {
	key      string
	response chan result
}
type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

type Memo struct{ requests chan request }

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.monitor(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	result := <-response
	return result.value, result.err
}

func (memo *Memo) monitor(f Func) {
	cache := make(map[string]*entry)

	for req := range memo.requests {
		_entry := cache[req.key]
		if _entry == nil {
			_entry = &entry{ready: make(chan struct{})}
			cache[req.key] = _entry
			go _entry.call(f, req.key)
		}
		go _entry.deliver(req.response)
	}

}

func (memo *Memo) Close() { close(memo.requests) }

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
