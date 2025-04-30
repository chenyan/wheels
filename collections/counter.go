package collections

import (
	"iter"
	"maps"
	"sort"
)

// Counter is a map of values to their counts.
type Counter[T comparable] struct {
	items map[T]int64
	sorted []T
	isSorted bool
}

func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{items: make(map[T]int64)}
}

func (c *Counter[T]) Add(value T) {
	c.items[value]++
	c.isSorted = false
}

func (c *Counter[T]) AddAll(values ...T) {
	for _, value := range values {
		c.items[value]++
	}
	c.isSorted = false
}

func (c *Counter[T]) Get(value T) int64 {
	return c.items[value]
}

func (c *Counter[T]) Sorted() []T {
	if !c.isSorted {
		keys := make([]T, 0, len(c.items))
		for k := range c.items {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return c.items[keys[i]] > c.items[keys[j]]
		})
		c.sorted = keys
		c.isSorted = true
	}
	return c.sorted
}

func (c *Counter[T]) Len() int {
	return len(c.items)
}


func (c *Counter[T]) MostCommon(n int) iter.Seq2[T, int64] {
	if n > c.Len() {
		n = c.Len()
	}
	ks := c.Sorted()[:n]
	return func(yield func(T, int64) bool) {
		for _, k := range ks {
			if !yield(k, c.items[k]) {
				return
			}
		}
	}
}

func (c *Counter[T]) LeastCommon(n int) iter.Seq2[T, int64] {
	if n > c.Len() {
		n = c.Len()
	}
	ks := c.Sorted()[c.Len()-n:]
	return func(yield func(T, int64) bool) {
		for i := 0; i < n; i++ {
			k := ks[n-i-1]
			if !yield(k, c.items[k]) {
				return
			}
		}
	}
}

func (c *Counter[T]) Delete(value T) {
	delete(c.items, value)
	c.isSorted = false
}

func (c *Counter[T]) Clear() {
	c.items = make(map[T]int64)
	c.sorted = nil
	c.isSorted = true
}

func (c *Counter[T]) Update(other *Counter[T]) {
	for k, v := range other.items {
		c.items[k] += v
	}
	c.isSorted = false
}

func (c *Counter[T]) Subtract(other *Counter[T]) {
	for k, v := range other.items {
		c.items[k] -= v
	}
	c.isSorted = false
}

func (c *Counter[T]) Copy() *Counter[T] {
	rv := NewCounter[T]()
	maps.Copy(rv.items, c.items)
	return rv
}

func (c *Counter[T]) Items() map[T]int64 {
	return c.items
}

func (c *Counter[T]) Keys() iter.Seq[T] {
	return maps.Keys(c.items)
}




