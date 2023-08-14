package renders

import (
	"bytes"
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	md := []byte("# Hello, world!")
	expected := []byte("<h1 id=\"hello-world\">Hello, world!</h1>\n")
	result := MarkdownToHTML(md)
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
