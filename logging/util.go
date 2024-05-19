package logging

import (
	"os"
	"strings"
)

// Shortpath returns the last n elements of the path p.
func Shortpath(p string, n uint) string {
	sep := string(os.PathSeparator)
	xs := strings.Split(p, sep)
	return strings.Join(xs[len(xs)-int(n):], sep)
}
