package main

import (
	"fmt"

	"github.com/mmcomp/cgo-example1/cgoexample"
)

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	cgoexample.Example(name)
}
