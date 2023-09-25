package server

import (
	"example/config"
	"example/handlers"
	"net/http"

	"github.com/google/wire"
)

type IServer interface {
	GetMux() *http.ServeMux
}

type Server struct {
	Cfg              config.IServerConfig
	TransformHandler handlers.ITransformHandler
}

var ServerSet = wire.NewSet(
	wire.Struct(new(Server), "*"),
	wire.Bind(new(IServer), new(*Server)),
)

func (s *Server) GetMux() *http.ServeMux {
	_ = s.Cfg
	mux := http.NewServeMux()
	mux.Handle("/", s.TransformHandler)

	return mux
}
