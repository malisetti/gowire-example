package handlers

import (
	"example/config"
	"example/hello"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/wire"
)

type ITransformHandler interface {
	http.Handler
}

type TransformHandler struct {
	cfg config.ITransformHandlerConfig
}

func NewTransformHandler(cfg config.ITransformHandlerConfig) *TransformHandler {
	return &TransformHandler{
		cfg: cfg,
	}
}

var TransformHandlerSet = wire.NewSet(
	NewTransformHandler,
	wire.Bind(new(ITransformHandler), new(*TransformHandler)),
)

func (th *TransformHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	transform, err := RequestToTransform(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
		return
	}
	sayer := hello.InitializeSayer(hello.TransformType(transform))

	sayer.Say(w, q.Get("message"))
}

func RequestToTransform(r *http.Request) (hello.TransformType, error) {
	q := r.URL.Query()
	transform := q.Get("transform")
	t, _ := strconv.Atoi(transform)
	tt := hello.TransformType(t)
	if tt < hello.ZeroTransform || tt > hello.UpperTransform {
		return hello.ZeroTransform, fmt.Errorf("invalid transform")
	}
	return tt, nil
}
