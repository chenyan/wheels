package collections

import (
	"iter"
)

// Set is a collection of unique items.
type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func NewSetFromSlice[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	set.Add(items...)
	return set
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {	
		s.items[item] = struct{}{}
	}
}

func (s *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s.items, item)
	}
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	return s.Intersection(other).IsEmpty()
}	

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.items))
	for item := range s.items {
		slice = append(slice, item)
	}
	return slice
}

func (s *Set[T]) ToSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s.items {
			if !yield(item) {
				return
			}
		}
	}
}

