package main

/*
#include "cgo.h"
*/
import (
	"C"
	"unsafe"
	"fmt"
)

func main() {
	var example *C.char = C.get_char_array()
	length := C.get_char_array_length()
	exampleSlice := (*[1 << 30]C.char)(unsafe.Pointer(example))[:length:length]
	fmt.Println( exampleSlice )
}