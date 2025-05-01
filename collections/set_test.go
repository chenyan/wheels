package collections

import (
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[string]()
	if s == nil {
		t.Error("NewSet() returned nil")
	}
	if !s.IsEmpty() {
		t.Error("NewSet() should be empty")
	}
}

func TestNewSetFromSlice(t *testing.T) {
	items := []string{"a", "b", "c", "a"}
	s := NewSetFromSlice(items)

	if s.Len() != 3 {
		t.Errorf("NewSetFromSlice() length = %d, want 3", s.Len())
	}
	if !s.Contains("a") || !s.Contains("b") || !s.Contains("c") {
		t.Error("NewSetFromSlice() missing items")
	}
}

func TestSet_Add(t *testing.T) {
	s := NewSet[string]()
	s.Add("a", "b", "a")

	if s.Len() != 2 {
		t.Errorf("Add() length = %d, want 2", s.Len())
	}
	if !s.Contains("a") || !s.Contains("b") {
		t.Error("Add() items not found in set")
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[string]()
	s.Add("a", "b", "c")
	s.Remove("a", "d") // 'd' doesn't exist

	if s.Len() != 2 {
		t.Errorf("Remove() length = %d, want 2", s.Len())
	}
	if s.Contains("a") {
		t.Error("Remove() 'a' still in set")
	}
	if !s.Contains("b") || !s.Contains("c") {
		t.Error("Remove() removed wrong items")
	}
}

func TestSet_Contains(t *testing.T) {
	s := NewSet[string]()
	s.Add("a")

	if !s.Contains("a") {
		t.Error("Contains() returned false for existing item")
	}
	if s.Contains("b") {
		t.Error("Contains() returned true for non-existing item")
	}
}

func TestSet_Len(t *testing.T) {
	s := NewSet[string]()
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}

	s.Add("a", "b", "a")
	if s.Len() != 2 {
		t.Errorf("Len() = %d, want 2", s.Len())
	}
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet[string]()
	if !s.IsEmpty() {
		t.Error("IsEmpty() returned false for empty set")
	}

	s.Add("a")
	if s.IsEmpty() {
		t.Error("IsEmpty() returned true for non-empty set")
	}
}

func TestSet_Clear(t *testing.T) {
	s := NewSet[string]()
	s.Add("a", "b", "c")
	s.Clear()

	if !s.IsEmpty() {
		t.Error("Clear() did not empty the set")
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b")
	s2.Add("b", "c")

	result := s1.Union(s2)
	if result.Len() != 3 {
		t.Errorf("Union() length = %d, want 3", result.Len())
	}
	if !result.Contains("a") || !result.Contains("b") || !result.Contains("c") {
		t.Error("Union() missing items")
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b")
	s2.Add("b", "c")

	result := s1.Intersection(s2)
	if result.Len() != 1 {
		t.Errorf("Intersection() length = %d, want 1", result.Len())
	}
	if !result.Contains("b") {
		t.Error("Intersection() missing common item")
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b", "c")
	s2.Add("b", "c", "d")

	result := s1.Difference(s2)
	if result.Len() != 1 {
		t.Errorf("Difference() length = %d, want 1", result.Len())
	}
	if !result.Contains("a") {
		t.Error("Difference() missing item")
	}
}

func TestSet_IsSubset(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b")
	s2.Add("a", "b", "c")

	if !s1.IsSubset(s2) {
		t.Error("IsSubset() returned false for valid subset")
	}
	if s2.IsSubset(s1) {
		t.Error("IsSubset() returned true for non-subset")
	}
}

func TestSet_IsSuperset(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b", "c")
	s2.Add("a", "b")

	if !s1.IsSuperset(s2) {
		t.Error("IsSuperset() returned false for valid superset")
	}
	if s2.IsSuperset(s1) {
		t.Error("IsSuperset() returned true for non-superset")
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()
	s1.Add("a", "b")
	s2.Add("c", "d")

	if !s1.IsDisjoint(s2) {
		t.Error("IsDisjoint() returned false for disjoint sets")
	}

	s2.Add("b")
	if s1.IsDisjoint(s2) {
		t.Error("IsDisjoint() returned true for non-disjoint sets")
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := NewSet[string]()
	s.Add("a", "b", "c")

	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("ToSlice() length = %d, want 3", len(slice))
	}

	// Sort for consistent comparison
	sort.Strings(slice)
	expected := []string{"a", "b", "c"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("ToSlice()[%d] = %s, want %s", i, slice[i], v)
		}
	}
}

func TestSet_ToSeq(t *testing.T) {
	s := NewSet[string]()
	s.Add("a", "b", "c")

	items := make(map[string]bool)
	for item := range s.ToSeq() {
		items[item] = true
	}

	if len(items) != 3 {
		t.Errorf("ToSeq() yielded %d items, want 3", len(items))
	}
	if !items["a"] || !items["b"] || !items["c"] {
		t.Error("ToSeq() missing items")
	}
} 