package files

import (
	"fmt"
	"testing"
)

func TestReadTextFile(t *testing.T) {
	bs, err := ReadTextFile("text.go")
	if err != nil {
		t.Fatalf("Failed to read text file: %v", err)
	}
	fmt.Println(bs)
}
