package main

import (
	"example/config"
	"example/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// hello.NewSayer(&hello.NewLine{}).Say(os.Stdout, hello.Message)
	// hello.InitializeSayer(hello.NewLineTransform).Say(os.Stdout, hello.Message)
	// parker := valet.InitializeValetParker(hello.ExactTransform, hello.Message)
	// _ = parker.Park(os.Stdout, valet.Car{})

	cfg := config.GetCfg()
	server := server.InitializeServer()
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), server.GetMux())
	if err != nil {
		log.Fatalln(err)
	}
}
