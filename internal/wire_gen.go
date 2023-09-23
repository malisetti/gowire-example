// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"example/hello"
	"example/internal/config"
	"example/valet"
	"github.com/google/wire"
	"io"
)

// Injectors from wire.go:

func InitializeSayer(w io.Writer, cfg hello.TransformerProviderConfig) *hello.ExactSayer {
	transformer := hello.NewTransform(cfg)
	exactSayer := hello.NewSayer(w, transformer)
	return exactSayer
}

func InitializeValetParker(w io.Writer, message string) *valet.ValetPark {
	configuration := config.GetCfg()
	transformer := hello.NewTransform(configuration)
	exactSayer := hello.NewSayer(w, transformer)
	valetPark := valet.NewValetParker(exactSayer, message)
	return valetPark
}

// wire.go:

var appSet = wire.NewSet(config.ConfigSet, hello.HelloSet, hello.TransformSet, valet.ValetSet)
