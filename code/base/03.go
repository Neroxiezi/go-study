package main

import "fmt"

//defer
// 在函数中 经常需要创建资源(比如 数据库链接 文件句柄 锁等) 为了在函数执行完毕后 及时的释放资源
// go 的设计者提供 defer(延时机制)

func main() {
	res := sum(10,20)
	fmt.Println("res=",res)
}

func sum(n1 int,n2 int) int {
	// 当执行到defer 时暂不执行 会将defer 后的语句压入 defer栈
	// 当函数执行完毕之后 按照先入后出的方式依次执行
	defer fmt.Println("ok n1",n1)
	defer fmt.Println("ok n2",n2)

	// defer 值拷贝
	n1++
	n2++
	res := n2 + n1
	fmt.Println("ok res", res)
	return res
}

// 细节说明
// 1. 当 go 执行到一个 defer 时, 不会立即执行 defer 后的语句 而是将 defer 后的语句压入到一个栈 中
// 2. 当函数执行完毕后 在从defer 栈中  依次从栈顶 去除语句执行 
// 3. 在defer 将语句放入到栈时,也会将相关的值拷贝同时入栈

