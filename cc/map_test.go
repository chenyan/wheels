package cc

import (
	"sync/atomic"
	"testing"
)

func TestMap_Load(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42

	// Test non-existent key
	if v, ok := m.Load(key); ok {
		t.Errorf("Load() = %v, %v, want _, false", v, ok)
	}

	// Test existing key
	m.Store(key, value)
	if v, ok := m.Load(key); !ok || v != value {
		t.Errorf("Load() = %v, %v, want %v, true", v, ok, value)
	}
}

func TestMap_Store(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42

	m.Store(key, value)
	if v, ok := m.Load(key); !ok || v != value {
		t.Errorf("Store() failed, Load() = %v, %v, want %v, true", v, ok, value)
	}
}

func TestMap_Delete(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42

	m.Store(key, value)
	m.Delete(key)
	if v, ok := m.Load(key); ok {
		t.Errorf("Delete() failed, Load() = %v, %v, want _, false", v, ok)
	}
}

func TestMap_Range(t *testing.T) {
	m := &Map[string, int]{}
	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range items {
		m.Store(k, v)
	}

	var count int32
	m.Range(func(key string, value int) bool {
		atomic.AddInt32(&count, 1)
		if items[key] != value {
			t.Errorf("Range() value mismatch for key %s: got %v, want %v", key, value, items[key])
		}
		return true
	})

	if count != int32(len(items)) {
		t.Errorf("Range() visited %d items, want %d", count, len(items))
	}
}

func TestMap_LoadOrStore(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42

	// Test storing new value
	if v, loaded := m.LoadOrStore(key, value); loaded {
		t.Errorf("LoadOrStore() = %v, %v, want _, false", v, loaded)
	}

	// Test loading existing value
	if v, loaded := m.LoadOrStore(key, 43); !loaded || v != value {
		t.Errorf("LoadOrStore() = %v, %v, want %v, true", v, loaded, value)
	}
}

func TestMap_CompareAndSwap(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	oldValue := 42
	newValue := 43

	// Test when key doesn't exist
	if m.CompareAndSwap(key, oldValue, newValue) {
		t.Error("CompareAndSwap() succeeded when key didn't exist")
	}

	// Test when old value doesn't match
	m.Store(key, oldValue)
	if m.CompareAndSwap(key, oldValue+1, newValue) {
		t.Error("CompareAndSwap() succeeded when old value didn't match")
	}

	// Test successful swap
	if !m.CompareAndSwap(key, oldValue, newValue) {
		t.Error("CompareAndSwap() failed when it should have succeeded")
	}
	if v, _ := m.Load(key); v != newValue {
		t.Errorf("CompareAndSwap() didn't update value, got %v, want %v", v, newValue)
	}
}

func TestMap_CompareAndDelete(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42

	// Test when key doesn't exist
	if m.CompareAndDelete(key, value) {
		t.Error("CompareAndDelete() succeeded when key didn't exist")
	}

	// Test when value doesn't match
	m.Store(key, value)
	if m.CompareAndDelete(key, value+1) {
		t.Error("CompareAndDelete() succeeded when value didn't match")
	}

	// Test successful delete
	if !m.CompareAndDelete(key, value) {
		t.Error("CompareAndDelete() failed when it should have succeeded")
	}
	if _, ok := m.Load(key); ok {
		t.Error("CompareAndDelete() didn't delete the key")
	}
}

func TestMap_Swap(t *testing.T) {
	m := &Map[string, int]{}
	key := "test"
	value := 42
	newValue := 43

	// Test swapping with non-existent key
	if v, loaded := m.Swap(key, value); loaded {
		t.Errorf("Swap() = %v, %v, want _, false", v, loaded)
	}

	// Test swapping with existing key
	m.Store(key, value)
	if v, loaded := m.Swap(key, newValue); !loaded || v != value {
		t.Errorf("Swap() = %v, %v, want %v, true", v, loaded, value)
	}
	if v, _ := m.Load(key); v != newValue {
		t.Errorf("Swap() didn't update value, got %v, want %v", v, newValue)
	}
}

func TestMap_Clear(t *testing.T) {
	m := &Map[string, int]{}
	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range items {
		m.Store(k, v)
	}

	m.Clear()
	if m.Len() != 0 {
		t.Errorf("Clear() failed, Len() = %d, want 0", m.Len())
	}
}

func TestMap_Len(t *testing.T) {
	m := &Map[string, int]{}
	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	if m.Len() != 0 {
		t.Errorf("Len() = %d, want 0", m.Len())
	}

	for k, v := range items {
		m.Store(k, v)
	}

	if m.Len() != len(items) {
		t.Errorf("Len() = %d, want %d", m.Len(), len(items))
	}
}

func TestMap_Keys(t *testing.T) {
	m := &Map[string, int]{}
	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range items {
		m.Store(k, v)
	}

	keys := m.Keys()
	if len(keys) != len(items) {
		t.Errorf("Keys() returned %d items, want %d", len(keys), len(items))
	}

	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	for k := range items {
		if !keySet[k] {
			t.Errorf("Keys() missing key %s", k)
		}
	}
}

func TestMap_Values(t *testing.T) {
	m := &Map[string, int]{}
	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range items {
		m.Store(k, v)
	}

	values := m.Values()
	if len(values) != len(items) {
		t.Errorf("Values() returned %d items, want %d", len(values), len(items))
	}

	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	for _, v := range items {
		if !valueSet[v] {
			t.Errorf("Values() missing value %d", v)
		}
	}
} 