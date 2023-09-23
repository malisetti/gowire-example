//go:build wireinject
// +build wireinject

package internal

import (
	"example/hello"
	"example/internal/config"
	"example/valet"
	"io"

	"github.com/google/wire"
)

var appSet = wire.NewSet(config.ConfigSet, hello.HelloSet, hello.TransformSet, valet.ValetSet)

func InitializeSayer(w io.Writer, t hello.TransformType) *hello.ExactSayer {
	wire.Build(appSet)
	return &hello.ExactSayer{}
}

func InitializeValetParker(w io.Writer, message string) *valet.ValetPark {
	wire.Build(appSet)
	return &valet.ValetPark{}
}
