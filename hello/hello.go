package hello

import (
	"io"

	"github.com/google/wire"
)

const Message = "Hello, World!"

type Sayer interface {
	Say(msg string) (n int, err error)
}

type ExactSayer struct {
	w io.Writer
	t Transformer
}

func NewSayer(w io.Writer, t Transformer) *ExactSayer {
	return &ExactSayer{
		w: w,
		t: t,
	}
}

var HelloSet = wire.NewSet(
	NewSayer,
	wire.Bind(new(Sayer), new(*ExactSayer)),
)

func (es *ExactSayer) Say(msg string) (int, error) {
	return es.write([]byte(es.t.Transform(msg)))
}

func (es *ExactSayer) write(p []byte) (int, error) {
	return es.w.Write(p)
}
