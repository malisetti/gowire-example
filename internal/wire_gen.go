// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"example/hello"
	"example/valet"
	"io"
)

// Injectors from wire.go:

func InitializeSayer(w io.Writer, cfg hello.TransformerProviderConfig) *hello.ExactSayer {
	transformer := hello.NewTransform(cfg)
	exactSayer := hello.NewSayer(w, transformer)
	return exactSayer
}

func InitializeValetParker(w io.Writer, cfg hello.TransformerProviderConfig, message string) *valet.ValetPark {
	transformer := hello.NewTransform(cfg)
	exactSayer := hello.NewSayer(w, transformer)
	valetPark := valet.NewValetParker(exactSayer, message)
	return valetPark
}