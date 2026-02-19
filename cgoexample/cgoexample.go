package cgoexample

// #include "hello.h"
import "C"

func Example() {
	C.hello()
}
