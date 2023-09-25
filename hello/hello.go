package hello

import (
	"io"

	"github.com/google/wire"
)

const Message = "Hello, World!"

type Sayer interface {
	Say(w io.Writer, msg string) (n int, err error)
}

type ExactSayer struct {
	t Transformer
}

func NewSayer(t Transformer) *ExactSayer {
	return &ExactSayer{
		t: t,
	}
}

var HelloSet = wire.NewSet(
	NewSayer,
	wire.Bind(new(Sayer), new(*ExactSayer)),
)

func (es *ExactSayer) Say(w io.Writer, msg string) (int, error) {
	return w.Write([]byte(es.t.Transform(msg)))
}
