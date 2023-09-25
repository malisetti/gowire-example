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
	hello.NewSayer(&hello.Exact{}).Say(os.Stdout, hello.Message)
	hello.InitializeSayer(hello.NewLineTransform).Say(os.Stdout, hello.Message)
	parker := valet.InitializeValetParker(hello.UpperTransform, hello.Message)
	_ = parker.Park(os.Stdout, valet.Car{})

	server, err := server.InitializeServer()
	if err != nil {
		panic(fmt.Sprintf("could not initialize server. failed with: %v", err))
	}
	cfg, _ := config.GetCfg()
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), server.GetMux())
	if err != nil {
		log.Fatalln(err)
	}
}
