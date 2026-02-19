package cgoexample

// #include "hello.h"
import "C"

func Example(name string) {
	C.hello(C.CString(name))
}
