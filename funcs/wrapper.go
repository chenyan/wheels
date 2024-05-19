package funcs

import (
	"log"
	"os"
	"runtime"
)

var (
	Errlogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
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
				Errlogger.Printf("recovered from panic: %v\n%s", r, buf[:n])
			} else {
				Errlogger.Printf("recovered from panic: %v", r)
			}
		}
	}()
	f()
}
