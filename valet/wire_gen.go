// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package valet

import (
	"example/hello"
)

// Injectors from wire.go:

func InitializeValetParker(tt hello.TransformType, message string) *ValetPark {
	transformer := hello.NewTransform(tt)
	exactSayer := hello.NewSayer(transformer)
	valetPark := NewValetParker(exactSayer, message)
	return valetPark
}
