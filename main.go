package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mmcomp/cgo-example1/cgoexample"
	// "github.com/mmcomp/cgo-example1/cgoexample"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// fmt.Println("Hello,", name)
	result := cgoexample.Example(name)
	fmt.Println(result)
}
