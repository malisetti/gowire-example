package hello

import (
	"io"
)

const Message = "Hello, World!"

type Sayer interface {
	Say(w io.Writer, msg string) (n int, err error)
}

type ExactSayer struct {
	w io.Writer
	t Transformer
}

func NewSayer(t Transformer) Sayer {
	return &ExactSayer{
		t: t,
	}
}

func (es *ExactSayer) Say(w io.Writer, msg string) (int, error) {
	es.w = w
	return w.Write([]byte(es.t.Transform(msg)))
}

func (es *ExactSayer) Write(p []byte) (int, error) {
	return es.w.Write(p)
}
