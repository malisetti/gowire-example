package config

import (
	"flag"
	"sync"

	"github.com/google/wire"
)

type TransformHandlerConfig interface{}

type IServerConfig interface {
	GetPort() uint
}

type IConfiguration interface {
	IServerConfig
	TransformHandlerConfig
}

type Configuration struct {
	TransformerProviderConfig
	ServerConfig
}

type TransformerProviderConfig struct{}

type ServerConfig struct {
	Port uint
}

func (cfg *ServerConfig) GetPort() uint {
	return cfg.Port
}

var cfg Configuration
var once sync.Once

const defaultPort = 3333

func GetCfg() (*Configuration, error) {
	var err error
	once.Do(func() {
		var port uint
		flag.UintVar(&port, "port", defaultPort, "Port to run. Default is 3333")
		flag.Parse()
		cfg.ServerConfig.Port = port
	})

	return &cfg, err
}

var ConfigSet = wire.NewSet(
	GetCfg,
	wire.Bind(new(IServerConfig), new(*Configuration)),
	wire.Bind(new(TransformHandlerConfig), new(*Configuration)),
)
