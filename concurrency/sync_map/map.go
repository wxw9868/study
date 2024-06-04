package sync_map

import "sync"

type Map struct {
	sync.Mutex
	data map[string]interface{}
}

func New() *Map {
	return &Map{data: make(map[string]interface{})}
}

func (m *Map) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func (m *Map) Get(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	if v, ok := m.data[key]; !ok {
		return nil
	} else {
		return v
	}
}

func (m *Map) Len() int {
	m.Lock()
	defer m.Unlock()
	return len(m.data)
}

func (m *Map) Print() map[string]interface{} {
	m.Lock()
	defer m.Unlock()
	return m.data
}

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}
