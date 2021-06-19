package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("当前系统是:%s\n", goos)
	path := os.Getenv("GOROOT")
	fmt.Printf("当前的路径是: %s\n", path)
}
