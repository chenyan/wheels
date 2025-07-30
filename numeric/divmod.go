package numeric

import "github.com/chenyan/wheels/types"

func Divmod[T types.Integer](a, b T) (T, T) {
	return a / b, a % b
}
