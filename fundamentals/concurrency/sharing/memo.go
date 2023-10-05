package sharing

import (
	"io"
	"net/http"
	"sync"
)

type Memo struct {
	f     Func
	cache map[string]Result
	mutex sync.Mutex
}

type Func func(key string) (interface{}, error)

type Result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]Result)}
}
func (m *Memo) Get(key string) (interface{}, error) {
	m.mutex.Lock()
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	m.mutex.Unlock()

	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
