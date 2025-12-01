package flow

import (
	"fmt"
	"os"
)

func GetenvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("env var %s is not set", key))
	}
	return val
}

func GetenvOr(key string, defval string) string {
	val := os.Getenv(key)
	if val == "" {
		return defval
	}
	return val
}
