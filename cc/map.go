package cc

import "sync"

// Map is a concurrent map implementation.
// It is a wrapper around sync.Map that provides a more convenient API.
type Map[K comparable, V any] struct {
	m sync.Map
}

// NewMap creates a new Map.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: sync.Map{},
	}
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, false
	}
	return v.(V), true
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.m.LoadOrStore(key, value)
	return v.(V), loaded
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, false
	}
	return v.(V), true
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) CompareAndDelete(key K, old V) bool {
	return m.m.CompareAndDelete(key, old)
}

func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := m.m.Swap(key, value)
	var zero V
	if !loaded {
		return zero, false
	}
	return v.(V), true
}

func (m *Map[K, V]) Clear() {
	m.m.Clear()
}

func (m *Map[K, V]) Len() int {
	len := 0
	m.Range(func(key K, value V) bool {
		len++
		return true
	})
	return len
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	m.Range(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	m.Range(func(key K, value V) bool {
		values = append(values, value)
		return true
	})
	return values
}
