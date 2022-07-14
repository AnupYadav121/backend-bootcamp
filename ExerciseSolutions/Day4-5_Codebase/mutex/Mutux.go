package mutex

import "sync"

type mutex struct {
	myMap sync.Map
}

var Mutex mutex

func (m *mutex) Lock(key string) bool {
	_, isPresent := m.myMap.Load(key)
	if isPresent == true {
		return false
	}

	m.myMap.Store(key, true)
	return true
}

func (m *mutex) UnLock(id string) {
	m.myMap.Delete(id)
}
