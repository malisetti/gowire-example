package hello

import (
	"fmt"
	"io"
)

const Message = "Hello, World!"

type Sayer interface {
	Say(w io.Writer, msg string)
}

var DefaultSayer NewLineSayer

func Say(w io.Writer, msg string) {
	DefaultSayer.Say(w, msg)
}

type NewLineSayer struct{}

func (nls NewLineSayer) Say(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

type CopySayer struct{}

func (cp CopySayer) Say(w io.Writer, msg string) {
	w.Write([]byte(msg))
}

type NoSayer struct{}

func (ns NoSayer) Say(w io.Writer, msg string) {
	_ = w
	_ = msg
}
