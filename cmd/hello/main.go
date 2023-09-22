package main

import (
	"example/hello"
	"fmt"
)

func main() {
	sayHello(hello.Message)
}

func sayHello(msg string) {
	fmt.Println(msg)
}
