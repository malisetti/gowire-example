//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"example/hello"
	"example/valet"

	"github.com/google/wire"
)

func InitializeValetParker(message string) valet.ValetPark {
	wire.Build(valet.NewValetParker, hello.NewSayer, hello.NewDefaultTransformer)
	return valet.ValetPark{}
}
