package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0,1,-2,3,4}
	fmt.Println(array)
	pointer :=&array[0]
	fmt.Println(*pointer," ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println(memoryAddress)
	for i := 0; i < len(array)-1; i++ {
        pointer = (*int)(unsafe.Pointer(memoryAddress))
        fmt.Print(*pointer, " ")
        memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
    }
}