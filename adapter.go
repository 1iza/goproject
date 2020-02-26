package syncAdapter

import "sync"

type adapterMap struct {
	srvToMap map[string]*sync.Map
	sync.RWMutex
}

func NewAdapter() *adapterMap {
	return &adapterMap{srvToMap: make(map[string]*sync.Map)}
}

func (a *adapterMap) IsExist(srvid string) (*sync.Map, bool) {
	a.RLock()
	defer a.RUnlock()
	if m, ok := a.srvToMap[srvid]; ok {
		return m, ok
	}
	return nil, false
}

func (a *adapterMap) New(srvid string, key uint64) chan []byte {
	c := make(chan []byte, 1)
	if m, ok := a.IsExist(srvid); ok {
		m.Store(key, c)
		return c
	}
	m := &sync.Map{}
	m.Store(key, c)
	a.Lock()
	defer a.Unlock()
	a.srvToMap[srvid] = m
	return c
}

func (a *adapterMap) Delete(srvid string, key uint64) {
	if m, ok := a.IsExist(srvid); ok {
		m.Delete(key)
	}
	return
}

func (a *adapterMap) Get(srvid string, key uint64) (chan []byte, bool) {
	m, ok := a.IsExist(srvid)
	if !ok {
		return nil, false
	}
	val, ok := m.Load(key)
	if !ok {
		return nil, false
	}
	if ch, ok := val.(chan []byte); ok {
		return ch, ok
	} else {
		return nil, false
	}
}
