package main

import (
	"fmt"
	"time"
)

func main() {
	// 闭包函数

	func (name string){
		fmt.Println("My name is ", name)
	}("王小二")

	// 匿名函数
	currentTime := func() {
		fmt.Println(time.Now())
	}
	currentTime()

}
