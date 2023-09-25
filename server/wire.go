//go:build wireinject
// +build wireinject

package server

import (
	"example/config"
	"example/handlers"

	"github.com/google/wire"
)

func InitializeServer() (*Server, error) {
	wire.Build(config.ConfigSet, handlers.HandlerSet, ServerSet)
	return &Server{}, nil
}
