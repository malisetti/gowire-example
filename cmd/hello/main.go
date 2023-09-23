package main

import (
	"example/hello"
	"example/internal"
	"example/valet"
	"os"
)

func main() {
	// hello.NewSayer(os.Stdout, &hello.NewLine{}).Say(hello.Message)
	internal.InitializeSayer(os.Stdout, hello.NewLineTransform).Say(hello.Message)
	parker := internal.InitializeValetParker(os.Stdout, hello.Message)
	_ = parker.Park(valet.Car{})
}
