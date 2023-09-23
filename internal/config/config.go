package config

import (
	"example/hello"
	"sync"

	"github.com/google/wire"
)

type Config interface {
	hello.TransformerProviderConfig
}

type Configuration struct {
	TransformerProviderType hello.TransformType
}

func (cfg *Configuration) GetTransformerProviderType() hello.TransformType {
	return cfg.TransformerProviderType
}

var cfg Configuration
var once sync.Once

func GetCfg() *Configuration {
	once.Do(func() {
		cfg.TransformerProviderType = hello.NewLineTransform
	})

	return &cfg
}

var ConfigSet = wire.NewSet(
	GetCfg,
	wire.Bind(new(Config), new(*Configuration)),
	wire.Bind(new(hello.TransformerProviderConfig), new(*Configuration)),
)
