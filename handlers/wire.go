//go:build wireinject
// +build wireinject

package handlers

import (
	"github.com/google/wire"
)

func InitializeTransformHandler() *TransformHandler {
	wire.Build(HandlerSet)
	return &TransformHandler{}
}
