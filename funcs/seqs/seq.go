package seqs

import (
	"iter"
)

// Apply applies the function f to each element of the slice ts in place and returns the resulting slice.
func Apply[T any](ts []T, f func(T) T) []T {
	for i, t := range ts {
		ts[i] = f(t)
	}
	return ts
}

// ToMap converts a slice to a map by applying the function f to each element of the slice.
func ToMap[T any](ts []T, f func(T) (string, T)) map[string]T {
	m := make(map[string]T, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

// Map applies the function f to each element of the slice ts and returns the new resulting slice.
func Map[T any](ts []T, f func(T) T) []T {
	rs := make([]T, len(ts))
	for i, t := range ts {
		rs[i] = f(t)
	}
	return rs
}

func ToSeq[T any](ts []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range ts {
			if !yield(t) {
				return
			}
		}
	}
}

// Zip zips two sequences together.
func Zip[A any, B any](as iter.Seq[A], bs iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		next1, stop1 := iter.Pull(as)
		next2, stop2 := iter.Pull(bs)
		defer stop1()
		defer stop2()
		for {
			a, ok1 := next1()
			b, ok2 := next2()
			if !ok1 && !ok2 {
				return
			}
			if ok1 != ok2 {
				panic("zip: sequences have different lengths")
			}
			if !yield(a, b) {
				return
			}
		}
	}
}
