package flow

import "log"

// PanicIf panic if cond is true
func PanicIf(cond bool, msg string) {
	if cond {
		log.Panic(msg)
	}
}

// PanicIfError panic if err is not nil
func PanicIfError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
