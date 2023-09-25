package config

import (
	"example/hello"
	"flag"
	"sync"

	"github.com/google/wire"
)

type IConfiguration interface {
	IServerConfig
	TransformHandlerConfig
}

type Configuration struct {
	TransformerProviderConfig struct {
		TransformerProviderType hello.TransformType
	}
	ServerConfig struct {
		Port uint
	}
}

func (cfg *Configuration) GetTransformerProviderType() hello.TransformType {
	return cfg.TransformerProviderConfig.TransformerProviderType
}

func (cfg *Configuration) GetPort() uint {
	return cfg.ServerConfig.Port
}

type TransformHandlerConfig interface {
	GetTransformerProviderType() hello.TransformType
}

type IServerConfig interface {
	GetPort() uint
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
		cfg.TransformerProviderConfig.TransformerProviderType = hello.TransformType(transformType)
		cfg.ServerConfig.Port = 3333
	})

	return &cfg
}

var ConfigSet = wire.NewSet(
	GetCfg,
	wire.Bind(new(IServerConfig), new(*Configuration)),
	wire.Bind(new(TransformHandlerConfig), new(*Configuration)),
)
