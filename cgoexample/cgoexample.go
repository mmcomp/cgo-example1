package cgoexample

// #include "hello.h"
import "C"
import "unsafe"

func Example(name string) string {
	result := C.hello(C.CString(name))
	ptr := uintptr(unsafe.Pointer(result))
	var bytes []byte
	for {
		b := *(*byte)(unsafe.Pointer(ptr))
		if b == 0 {
			break
		}
		bytes = append(bytes, b)
		ptr++
	}
	goStr := string(bytes)
	return goStr
}
