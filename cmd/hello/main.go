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
	cfg := config.GetCfg()

	hello.NewSayer(&hello.NewLine{}).Say(os.Stdout, hello.Message)
	hello.InitializeSayer(hello.NewLineTransform).Say(os.Stdout, hello.Message)
	parker := valet.InitializeValetParker(hello.ExactTransform, hello.Message)
	_ = parker.Park(os.Stdout, valet.Car{})

	server := server.InitializeServer()
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), server.GetMux())
	if err != nil {
		log.Fatalln(err)
	}
}
