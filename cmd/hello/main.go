package main

import (
	"example/hello"
	"example/internal"
	"example/internal/config"
	"example/valet"
	"os"
)

func main() {
	// hello.NewSayer(os.Stdout, &hello.NewLine{}).Say(hello.Message)
	cfg := config.GetCfg()
	internal.InitializeSayer(os.Stdout, cfg).Say(hello.Message)
	parker := internal.InitializeValetParker(os.Stdout, cfg, hello.Message)
	_ = parker.Park(valet.Car{})
}
