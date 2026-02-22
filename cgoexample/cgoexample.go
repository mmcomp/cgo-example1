package cgoexample

// #include "hello.h"
import "C"
import "unsafe"

func Example(name string) string {
	result := C.hello(C.CString(name))
	ln := len(name) + 12
	ptr := unsafe.Pointer(result)
	var bytes []byte
	for i := 0; i < ln; i++ {
		b := *(*byte)(unsafe.Pointer(ptr))
		bytes = append(bytes, b)
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}
	return string(bytes)
	// ptr := uintptr(unsafe.Pointer(result))
	// var bytes []byte
	// for {
	// 	b := *(*byte)(unsafe.Pointer(ptr))
	// 	if b == 0 {
	// 		break
	// 	}
	// 	bytes = append(bytes, b)
	// 	ptr++
	// }
	// goStr := string(bytes)
	// return goStr
}
