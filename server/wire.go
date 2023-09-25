//go:build wireinject
// +build wireinject

package server

import (
	"example/config"
	"example/handlers"

	"github.com/google/wire"
)

func InitializeServer(cfg config.IConfiguration) *Server {
	wire.Build(config.ConfigSet, handlers.HandlerSet, ServerSet)
	return &Server{}
}
