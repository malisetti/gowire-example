package main

import (
	"example/hello"
	"os"
)

func main() {
	hello.Say(os.Stdout, hello.Message)
}
