package main

import "fmt"

func proc(input string, processor func(str string)) {
	//调用匿名函数
	processor(input)
}

func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}
	return func() int {
		initial++
		// 返回当前计数
		return initial
	}
}

func main() {

	// 计数器1
	c1 := createCounter(1)
	fmt.Println(c1())
	fmt.Println(c1())

	// 计数器
	c2 := createCounter(10)
	fmt.Println(c2())
	fmt.Println(c2())

	proc("王小二", func(str string) {
		for _, v := range str {
			fmt.Printf("%c\n", v)
		}
	})
}
