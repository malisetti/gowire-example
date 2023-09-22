package hello

import (
	"fmt"
	"io"
)

const Message = "Hello, World!"

type Sayer interface {
	Say(w io.Writer, msg string) (n int, err error)
}

var defaultSayer NewLineSayer

func Say(w io.Writer, msg string) {
	defaultSayer.Say(w, msg)
}

type NewLineSayer struct{}

func (nls NewLineSayer) Say(w io.Writer, msg string) (int, error) {
	return fmt.Fprintln(w, msg)
}

type CopySayer struct{}

func (cp CopySayer) Say(w io.Writer, msg string) (int, error) {
	return w.Write([]byte(msg))
}

type NoSayer struct{}

func (ns NoSayer) Say(w io.Writer, msg string) (n int, err error) {
	_ = w
	_ = msg
	return
}
