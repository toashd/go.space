package main

import (
		"hello"
		"space"
		"os"
)

func main() {
		hello.SayHello(os.Stdout)
		space.SaySpace(os.Stdout)
}

