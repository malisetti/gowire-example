package main

import (
	"example/hello"
	"os"
)

func main() {
	hello.NewSayer(hello.NewLine{}).Say(os.Stdout, hello.Message)
}
