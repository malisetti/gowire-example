// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handlers

import (
	"example/config"
)

// Injectors from wire.go:

func InitializeTransformHandler(cfg config.ITransformHandlerConfig) *TransformHandler {
	transformHandler := NewTransformHandler(cfg)
	return transformHandler
}
