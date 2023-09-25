package handlers

import (
	"example/config"
	"example/hello"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/wire"
)

type ITransformHandler http.Handler

type TransformHandler struct {
	cfg config.TransformHandlerConfig
}

func NewTransformHandler(cfg config.TransformHandlerConfig) *TransformHandler {
	return &TransformHandler{
		cfg: cfg,
	}
}

var HandlerSet = wire.NewSet(
	NewTransformHandler,
	wire.Bind(new(ITransformHandler), new(*TransformHandler)),
)

func (th *TransformHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	transform := q.Get("transform")
	message := q.Get("message")
	t, _ := strconv.Atoi(transform)
	if hello.TransformType(t) < hello.ZeroTransform || hello.TransformType(t) > hello.UpperTransform {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid transform\n")
		return
	}

	sayer := hello.InitializeSayer(hello.TransformType(t))
	sayer.Say(w, message)
}
