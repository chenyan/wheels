package seqs

// InitSlice initializes a slice with a given length and fill value
func InitSlice[T any](n int, fill T) []T {
	slice := make([]T, n)
	for i := range slice {
		slice[i] = fill
	}
	return slice
}

// InitSliceWith initializes a slice with a given length and fill function
func InitSliceWith[T any](n int, fill func(i int) T) []T {
	slice := make([]T, n)
	for i := range slice {
		slice[i] = fill(i)
	}
	return slice
}
