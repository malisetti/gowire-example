//go:build wireinject
// +build wireinject

package handlers

import (
	"example/config"

	"github.com/google/wire"
)

func InitializeTransformHandler(cfg config.ITransformHandlerConfig) *TransformHandler {
	wire.Build(HandlerSet)
	return &TransformHandler{}
}
