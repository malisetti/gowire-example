package server

import (
	"example/hello"
	"example/internal"
	"example/internal/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/wire"
)

type IServer interface {
	GetMux() *http.ServeMux
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var ServerSet = wire.NewSet(
	NewServer,
	wire.Bind(new(IServer), new(*Server)),
)

func (s *Server) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TransformServer)

	return mux
}

func TransformServer(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	transform := q.Get("transform")
	message := q.Get("message")

	t, _ := strconv.Atoi(transform)
	if hello.TransformType(t) < hello.ZeroTransform || hello.TransformType(t) > hello.UpperTransform {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid transform\n")
		return
	}
	cfg := config.TransformerProviderConfig{
		TransformerProviderType: hello.TransformType(t),
	}

	sayer := internal.InitializeSayer(w, &cfg)
	sayer.Say(message)
}
