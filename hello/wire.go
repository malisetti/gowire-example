//go:build wireinject
// +build wireinject

package hello

import (
	"github.com/google/wire"
)

func InitializeSayer(tt TransformType) *ExactSayer {
	wire.Build(HelloSet, TransformSet)
	return &ExactSayer{}
}
