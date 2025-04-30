package collections

import (
	"testing"
)

func TestNewCounter(t *testing.T) {
	c := NewCounter[string]()
	if c == nil {
		t.Error("NewCounter() returned nil")
	}
	if c.Len() != 0 {
		t.Errorf("NewCounter() length = %d, want 0", c.Len())
	}
}

func TestCounter_Add(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a", "b")

	if c.Get("a") != 2 {
		t.Errorf("Add() count for 'a' = %d, want 2", c.Get("a"))
	}
	if c.Get("b") != 1 {
		t.Errorf("Add() count for 'b' = %d, want 1", c.Get("b"))
	}
}

func TestCounter_Get(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a")

	if c.Get("a") != 2 {
		t.Errorf("Get() = %d, want 2", c.Get("a"))
	}
	if c.Get("b") != 0 {
		t.Errorf("Get() for non-existent key = %d, want 0", c.Get("b"))
	}
}

func TestCounter_Sorted(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("b", "a", "a", "c", "a")

	sorted := c.Sorted()
	expected := []string{"a", "b", "c"}
	if len(sorted) != len(expected) {
		t.Errorf("Sorted() length = %d, want %d", len(sorted), len(expected))
	}
	if sorted[0] != "a" {
		t.Errorf("Sorted()[0] = %s, want 'a'", sorted[0])
	}
}

func TestCounter_Len(t *testing.T) {
	c := NewCounter[string]()
	if c.Len() != 0 {
		t.Errorf("Len() = %d, want 0", c.Len())
	}

	c.AddAll("a", "b")
	if c.Len() != 2 {
		t.Errorf("Len() = %d, want 2", c.Len())
	}
}

func TestCounter_MostCommon(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a", "b", "c", "c", "c")

	var results []string
	for k, _ := range c.MostCommon(2) {
		results = append(results, k)
	}

	if len(results) != 2 {
		t.Errorf("MostCommon() returned %d items, want 2", len(results))
	}
	if results[0] != "c" {
		t.Errorf("MostCommon()[0] = %s, want 'c'", results[0])
	}
	if results[1] != "a" {
		t.Errorf("MostCommon()[1] = %s, want 'a'", results[1])
	}
}

func TestCounter_LeastCommon(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a", "b", "c", "c", "c")

	var results []string
	for k, _ := range c.LeastCommon(2) {
		results = append(results, k)
	}

	if len(results) != 2 {
		t.Errorf("LeastCommon() returned %d items, want 2", len(results))
	}
	if results[0] != "b" {
		t.Errorf("LeastCommon()[0] = %s, want 'b'", results[0])
	}
	if results[1] != "a" {
		t.Errorf("LeastCommon()[1] = %s, want 'a'", results[1])
	}
}

func TestCounter_Delete(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a", "b")

	c.Delete("a")
	if c.Get("a") != 0 {
		t.Errorf("Delete() failed, count for 'a' = %d, want 0", c.Get("a"))
	}
	if c.Get("b") != 1 {
		t.Errorf("Delete() affected wrong key, count for 'b' = %d, want 1", c.Get("b"))
	}
}

func TestCounter_Clear(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "b", "c")

	c.Clear()
	if c.Len() != 0 {
		t.Errorf("Clear() failed, Len() = %d, want 0", c.Len())
	}
	if c.Get("a") != 0 {
		t.Errorf("Clear() failed, count for 'a' = %d, want 0", c.Get("a"))
	}
}

func TestCounter_Update(t *testing.T) {
	c1 := NewCounter[string]()
	c2 := NewCounter[string]()
	c1.AddAll("a", "a")
	c2.AddAll("a", "b")

	c1.Update(c2)
	if c1.Get("a") != 3 {
		t.Errorf("Update() count for 'a' = %d, want 3", c1.Get("a"))
	}
	if c1.Get("b") != 1 {
		t.Errorf("Update() count for 'b' = %d, want 1", c1.Get("b"))
	}
}

func TestCounter_Subtract(t *testing.T) {
	c1 := NewCounter[string]()
	c2 := NewCounter[string]()
	c1.AddAll("a", "a", "a")
	c2.AddAll("a", "b")

	c1.Subtract(c2)
	if c1.Get("a") != 2 {
		t.Errorf("Subtract() count for 'a' = %d, want 2", c1.Get("a"))
	}
	if c1.Get("b") != -1 {
		t.Errorf("Subtract() count for 'b' = %d, want -1", c1.Get("b"))
	}
}

func TestCounter_Copy(t *testing.T) {
	c1 := NewCounter[string]()
	c1.AddAll("a", "a", "b")

	c2 := c1.Copy()
	if c2.Len() != c1.Len() {
		t.Errorf("Copy() length = %d, want %d", c2.Len(), c1.Len())
	}
	if c2.Get("a") != c1.Get("a") {
		t.Errorf("Copy() count for 'a' = %d, want %d", c2.Get("a"), c1.Get("a"))
	}

	// Test that copy is independent
	c2.Add("a")
	if c1.Get("a") == c2.Get("a") {
		t.Error("Copy() is not independent")
	}
}

func TestCounter_Items(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "a", "b")

	items := c.Items()
	if len(items) != 2 {
		t.Errorf("Items() length = %d, want 2", len(items))
	}
	if items["a"] != 2 {
		t.Errorf("Items() count for 'a' = %d, want 2", items["a"])
	}
	if items["b"] != 1 {
		t.Errorf("Items() count for 'b' = %d, want 1", items["b"])
	}
}

func TestCounter_Keys(t *testing.T) {
	c := NewCounter[string]()
	c.AddAll("a", "b", "c")

	keys := make(map[string]bool)
	for k := range c.Keys() {
		keys[k] = true
	}

	if len(keys) != 3 {
		t.Errorf("Keys() returned %d items, want 3", len(keys))
	}
	if !keys["a"] || !keys["b"] || !keys["c"] {
		t.Error("Keys() missing some keys")
	}
} 