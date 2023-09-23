package main

import (
	"example/hello"
	"example/internal"
	"example/internal/config"
	"example/server"
	"example/valet"
	"log"
	"net/http"
	"os"
)

func main() {
	// hello.NewSayer(os.Stdout, &hello.NewLine{}).Say(hello.Message)
	cfg := config.GetCfg()
	internal.InitializeSayer(os.Stdout, cfg).Say(hello.Message)
	parker := internal.InitializeValetParker(os.Stdout, cfg, hello.Message)
	_ = parker.Park(valet.Car{})

	s := server.NewServer()
	err := http.ListenAndServe(":3333", s.GetMux())
	if err != nil {
		log.Fatalln(err)
	}
}
