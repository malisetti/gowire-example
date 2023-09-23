//go:build wireinject
// +build wireinject

package internal

import (
	"example/hello"
	"example/valet"
	"io"

	"github.com/google/wire"
)

func InitializeSayer(w io.Writer, cfg hello.TransformerProviderConfig) *hello.ExactSayer {
	wire.Build(hello.HelloSet, hello.TransformSet)
	return &hello.ExactSayer{}
}

func InitializeValetParker(w io.Writer, cfg hello.TransformerProviderConfig, message string) *valet.ValetPark {
	wire.Build(hello.HelloSet, hello.TransformSet, valet.ValetSet)
	return &valet.ValetPark{}
}
