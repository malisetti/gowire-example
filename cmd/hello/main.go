package main

import (
	"example/hello"
	"example/valet"
	"os"
)

func main() {
	hello.NewSayer(hello.NewLine{}).Say(os.Stdout, hello.Message)

	parker := InitializeValetParker(hello.Message)
	_ = parker.Park(valet.Car{})
}
