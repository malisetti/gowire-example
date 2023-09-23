package valet

import (
	"example/hello"

	"github.com/google/wire"
)

type Vehicle interface {
	Drive()
}

type Car struct{}

func (Car) Drive() {}

type Valet interface {
	Park(v Vehicle) error
}

type ValetPark struct {
	helloSayer hello.Sayer
	message    string
}

func NewValetParker(sayer hello.Sayer, message string) *ValetPark {
	return &ValetPark{
		helloSayer: sayer,
		message:    message,
	}
}

var ValetSet = wire.NewSet(
	NewValetParker,
	wire.Bind(new(Valet), new(*ValetPark)),
)

func (vp *ValetPark) Park(v Vehicle) error {
	_, err := vp.helloSayer.Say(vp.message)
	if err != nil {
		return err
	}
	v.Drive()

	return nil
}
