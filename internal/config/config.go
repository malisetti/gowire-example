package config

import (
	"example/hello"
	"flag"
	"sync"

	"github.com/google/wire"
)

type Config interface {
	hello.TransformerProviderConfig
}

type TransformerProviderConfig struct {
	TransformerProviderType hello.TransformType
}

type Configuration struct {
	TransformerProviderConfig
}

func (cfg *TransformerProviderConfig) GetTransformerProviderType() hello.TransformType {
	return cfg.TransformerProviderType
}

var cfg Configuration
var once sync.Once

func GetCfg() *Configuration {
	once.Do(func() {
		var transformType int
		flag.IntVar(&transformType, "transform", int(hello.ExactTransform), "ZeroTransform 0, NewLineTransform 1, ExactTransform 2")
		flag.Parse()
		if transformType < int(hello.ZeroTransform) || transformType > int(hello.ExactTransform) {
			panic("invalid transform")
		}
		cfg.TransformerProviderType = hello.TransformType(transformType)
	})

	return &cfg
}

var ConfigSet = wire.NewSet(
	GetCfg,
	wire.Bind(new(hello.TransformerProviderConfig), new(*Configuration)),
)
