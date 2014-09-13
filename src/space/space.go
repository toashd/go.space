package space

import (
		"fmt"
		"io"
)

func SaySpace(out io.Writer) {
		fmt.Fprintf(out, "hello from space\n")
}

