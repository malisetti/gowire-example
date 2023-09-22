package main

import (
	"example/hello"
	"fmt"
)

func main() {
	sayHello()
}

func sayHello() {
	fmt.Println(hello.Message)
}
