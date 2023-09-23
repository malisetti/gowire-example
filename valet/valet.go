package valet

import (
	"example/hello"
	"os"
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
}

func NewValetParker(sayer hello.Sayer) ValetPark {
	return ValetPark{
		helloSayer: sayer,
	}
}

func (vp ValetPark) Park(v Vehicle) error {
	_, err := vp.helloSayer.Say(os.Stdout, hello.Message)
	if err != nil {
		return err
	}
	v.Drive()

	return nil
}
