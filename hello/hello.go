package hello

import (
	"fmt"
	"os"
)

const Message = "Hello, World!"

func Say(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}
