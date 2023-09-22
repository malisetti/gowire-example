package hello

import (
	"bytes"
	"strings"
	"testing"
)

func TestSay(t *testing.T) {
	var buf bytes.Buffer
	Say(&buf, Message)

	output := strings.TrimRight(buf.String(), "\n")
	if output != Message {
		t.Fatalf("output: '%s' not equals message: '%s'", output, Message)
	}
}
