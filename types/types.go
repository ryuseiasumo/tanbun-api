package types

import (
	"sync"
)

type SafeMap struct {
	v   map[string]string
	mux sync.Mutex
}

func (m *SafeMap) Get(key string) string {
	m.mux.Lock()
	value, ok := m.v[key]
	m.mux.Unlock()
	if ok {
		return value
	}
	return ""
}

func (m *SafeMap) Set(key string, value string) {
	m.mux.Lock()
	m.v[key] = value
	m.mux.Unlock()
}

func (m *SafeMap) ExistKey(key string) bool {
	m.mux.Lock()
	_, ok := m.v[key]
	m.mux.Unlock()
	return ok
}

func (m *SafeMap) RemoveByKey(key string) {
	m.mux.Lock()
	delete(m.v, key)
	m.mux.Unlock()
}

func (m *SafeMap) Init() {
	m.v = make(map[string]string)
}

// apierr.Code = 404
// apierr.Message = "Not Found"

type APIError struct {
	Code int
	Message string
}
