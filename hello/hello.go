package hello

import (
	"fmt"
	"io"
)

const Message = "Hello, World!"

func Say(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}
