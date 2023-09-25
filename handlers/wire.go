//go:build wireinject
// +build wireinject

package handlers

import (
	"example/config"

	"github.com/google/wire"
)

func InitializeTransformHandler(cfg config.TransformHandlerConfig) *TransformHandler {
	wire.Build(NewTransformHandler)
	return &TransformHandler{}
}
