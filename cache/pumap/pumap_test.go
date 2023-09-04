// BEGIN: 7c8c5d8d7b5c
package pumap

import (
	"fmt"
	"testing"
	"time"
)

func TestPUMap(t *testing.T) {
	m := NewPUMap[int, string](time.Second, func() map[int]string {
		return map[int]string{
			1: "one",
			2: "two",
			3: "three",
		}
	})
	m.Start()
	defer m.Stop()

	time.Sleep(2 * time.Second)

	if v, ok := m.Get(1); !ok || v != "one" {
		t.Errorf("expected value 'one' for key 1, got %v", v)
	}

	m.Put(4, "four")

	if v, ok := m.Get(4); !ok || v != "four" {
		t.Errorf("expected value 'four' for key 4, got %v", v)
	}

	m.Remove(2)

	if v, ok := m.Get(2); ok {
		t.Errorf("expected key 2 to be removed, got value %v", v)
	}

	fmt.Println("PUMap tests passed")
}

// END: 7c8c5d8d7b5c
