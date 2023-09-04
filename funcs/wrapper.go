package funcs

import (
	"log"
	"os"
	"runtime"
)

var (
	errlogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	ShowStack bool
	StackSize = 2048
)

// F is a wrapper for func() to recover from panic
func F(f func()) {
	defer func() {
		if r := recover(); r != nil {
			if ShowStack {
				buf := make([]byte, StackSize)
				n := runtime.Stack(buf, false)
				errlogger.Printf("recovered from panic: %v\n%s", r, buf[:n])
			} else {
				errlogger.Printf("recovered from panic: %v", r)
			}
		}
	}()
	f()
}
