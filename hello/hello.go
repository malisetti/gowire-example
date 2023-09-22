package hello

import (
	"io"
)

const Message = "Hello, World!"

type Sayer interface {
	io.Writer
	Say(w io.Writer, msg string) (n int, err error)
}

type ExactSayer struct {
	io.Writer
	Transformer
}

func NewSayer(f Transformer) Sayer {
	return &ExactSayer{
		Transformer: f,
	}
}

func (es *ExactSayer) Say(w io.Writer, msg string) (int, error) {
	es.Writer = w
	return w.Write([]byte(es.Transformer.Transform(msg)))
}

func (es *ExactSayer) Write(p []byte) (int, error) {
	return es.Writer.Write(p)
}
