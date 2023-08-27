package seq

// Apply applies the function f to each element of the slice ts and returns the resulting slice.
func Apply[T any](ts []T, f func(T) T) []T {
	for i, t := range ts {
		ts[i] = f(t)
	}
	return ts
}
