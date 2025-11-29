package numeric

import "cmp"

// Max returns the maximum value in the slice.
func Max[E cmp.Ordered](xs ...E) E {
	if len(xs) == 0 {
		panic("numeric.Max: empty slice")
	}
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

// Min returns the minimum value in the slice.
func Min[E cmp.Ordered](xs ...E) E {
	if len(xs) == 0 {
		panic("numeric.Min: empty slice")
	}
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}
