package main

import (
	"example/config"
	"example/hello"
	"example/server"
	"example/valet"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// hello.NewSayer(os.Stdout, &hello.NewLine{}).Say(hello.Message)
	cfg := config.GetCfg()
	hello.InitializeSayer(hello.ExactTransform).Say(os.Stdout, hello.Message)
	parker := valet.InitializeValetParker(hello.ExactTransform, hello.Message)
	_ = parker.Park(os.Stdout, valet.Car{})

	server := server.InitializeServer(cfg)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), server.GetMux())
	if err != nil {
		log.Fatalln(err)
	}
}
