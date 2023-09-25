//go:build wireinject
// +build wireinject

package valet

import (
	"example/hello"

	"github.com/google/wire"
)

func InitializeValetParker(tt hello.TransformType, message string) *ValetPark {
	wire.Build(hello.HelloSet, hello.TransformSet, ValetSet)
	return &ValetPark{}
}
