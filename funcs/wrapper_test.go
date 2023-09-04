// BEGIN: 8f7e2c1d3b4a
package funcs_test

import (
	"testing"

	"github.com/chenyan/wheels/funcs"
)

func TestF(t *testing.T) {
	// Test that F recovers from panic
	funcs.ShowStack = true
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("F did not recover from panic: %v", r)
			}
		}()
		funcs.F(func() {
			panic("test panic")
		})
	}()
}

// END: 8f7e2c1d3b4a
