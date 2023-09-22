package hello

import (
	"io"
	"os"
)

const Message = "Hello, World!"

type Sayer interface {
	io.Writer
	Say(w io.Writer, msg string) (n int, err error)
}

type NewLineSayer struct {
	io.Writer
}

type ExactSayer struct {
	io.Writer
}

type NoSayer struct {
	io.Writer
}

var defaultSayer NewLineSayer

func Say(msg string) {
	defaultSayer.Say(os.Stdout, msg)
}

func (nls *NewLineSayer) Say(w io.Writer, msg string) (int, error) {
	nls.Writer = w
	return nls.Write([]byte(msg + "\n"))
}

func (nls *NewLineSayer) Write(p []byte) (n int, err error) {
	return nls.Writer.Write(p)
}

func (es *ExactSayer) Say(w io.Writer, msg string) (int, error) {
	es.Writer = w
	return w.Write([]byte(msg))
}

func (es *ExactSayer) Write(p []byte) (n int, err error) {
	return es.Writer.Write(p)
}

func (ns *NoSayer) Say(w io.Writer, msg string) (n int, err error) {
	ns.Writer = w
	return ns.Write([]byte(msg))
}

func (ns *NoSayer) Write(p []byte) (n int, err error) {
	return
}
