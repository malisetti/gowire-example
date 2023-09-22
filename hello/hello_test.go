package hello_test

import (
	"bytes"
	"example/hello"
	"strings"
	"testing"
)

func TestSay(t *testing.T) {
	var buf bytes.Buffer
	hello.Say(&buf, hello.Message)

	output := strings.TrimRight(buf.String(), "\n")
	if output != hello.Message {
		t.Fatalf("output: '%s' not equals message: '%s'", output, hello.Message)
	}

	buf.Reset()

	var sayer hello.Sayer
	sayer = hello.NewLineSayer{}
	_, err := sayer.Say(&buf, hello.Message)
	if err != nil {
		t.Fatalf("default sayer could not say %s", hello.Message)
	}

	output = strings.TrimRight(buf.String(), "\n")
	if output != hello.Message {
		t.Fatalf("output: '%s' not equals message: '%s'", output, hello.Message)
	}

	buf.Reset()

	sayer = hello.CopySayer{}
	_, err = sayer.Say(&buf, hello.Message)
	if err != nil {
		t.Fatalf("copy sayer could not say %s", hello.Message)
	}
	output = buf.String()
	if output != hello.Message {
		t.Fatalf("output: '%s' not equals message: '%s'", output, hello.Message)
	}

	buf.Reset()

	sayer = hello.NoSayer{}
	n, err := sayer.Say(&buf, hello.Message)
	if err != nil {
		t.Fatalf("no sayer could not say %s", hello.Message)
	}
	if n > 0 || len(buf.String()) > 0 {
		t.Fatalf("output is not empty")
	}
}
