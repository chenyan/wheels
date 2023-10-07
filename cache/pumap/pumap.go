package pumap

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/chenyan/wheels/funcs"
)

// PUMap is a concurrent map with a periodic update mechanism.
// It is implemented as a struct with a map, a read-write mutex, a fetcher function, and an interval.
// The map keys must be comparable and the values can be of any type.
// The fetcher function is used to update the map periodically.
// The running field is an atomic boolean that indicates whether the update mechanism is running or not.
type PUMap[K comparable, V any] struct {
	sync.RWMutex
	m        map[K]V
	interval time.Duration
	fetcher  func() map[K]V
	running  atomic.Bool
}

// NewPUMap creates a new PUMap
func NewPUMap[K comparable, V any](interval time.Duration, fetcher func() map[K]V) *PUMap[K, V] {
	return &PUMap[K, V]{m: make(map[K]V), interval: interval, fetcher: fetcher}
}

func (m *PUMap[K, V]) Put(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *PUMap[K, V]) Get(k K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *PUMap[K, V]) Remove(k K) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *PUMap[K, V]) Start() {
	if m.running.Load() {
		return
	}
	m.running.Store(true)
	go func() {
		for m.running.Load() {
			m.Lock()
			funcs.F(func() { m.m = m.fetcher() })
			m.Unlock()
			time.Sleep(m.interval)
		}
	}()
}

func (m *PUMap[K, V]) Stop() {
	m.running.Store(false)
}
