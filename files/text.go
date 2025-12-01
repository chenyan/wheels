package files

import (
	"bytes"
	"errors"
	"os"
)

func ReadTextFile(path string) (string, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	// check if the file is binary
	var probeSize = 1024
	if len(bs) < probeSize {
		probeSize = len(bs)
	}
	buf := bs[:probeSize]
	isBinary := bytes.Contains(buf, []byte{0})
	if isBinary {
		return "", errors.New("file is binary")
	}
	return string(bs), nil
}
