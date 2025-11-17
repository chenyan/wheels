package seqs

// Apply applies the function f to each element of the slice ts and returns the resulting slice.
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
