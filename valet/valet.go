package valet

import (
	"example/hello"
	"io"

	"github.com/google/wire"
)

type Vehicle interface {
	Drive()
}

type Car struct{}

func (Car) Drive() {}

type Valet interface {
	Park(w io.Writer, v Vehicle) error
}

type ValetPark struct {
	helloSayer hello.Sayer
	message    string
}

var ValetSet = wire.NewSet(
	NewValetParker,
	wire.Bind(new(Valet), new(*ValetPark)),
)

func NewValetParker(sayer hello.Sayer, message string) *ValetPark {
	return &ValetPark{
		helloSayer: sayer,
		message:    message,
	}
}

func (vp *ValetPark) Park(w io.Writer, v Vehicle) error {
	_, err := vp.helloSayer.Say(w, vp.message)
	if err != nil {
		return err
	}
	v.Drive()

	return nil
}
