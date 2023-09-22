package main

import (
	"example/hello"
	"fmt"
)

func main() {
	say(hello.Message)
}

func say(msg string) {
	fmt.Println(msg)
}
