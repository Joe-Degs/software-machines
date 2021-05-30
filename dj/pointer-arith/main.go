package main

import (
	"fmt"
	"unsafe"
)

func main() {
	buf := []byte{1, 2, 3, 4, 5, 6}
	fmt.Println(buf)

	for i := 0; i < len(buf); i++ {
		fmt.Println(
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + unsafe.Sizeof(buf[0])*uintptr(i))),
		)
	}
}
