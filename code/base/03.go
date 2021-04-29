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

//a. defer 最主要的价值在 当函数执行完毕后 可以及时的释放函数创建的资源
//b. 在go 编程中通常做法是  创建资源后 比如(打开了文件, 获取了数据库的链接 或者是 锁资源) 可以执行 defer file.Close() defer connect.Close()
//c. 在 defer 后, 可以继续使用创建资源
//d. 当函数完毕后 系统依次从 defer 栈中 取出语句 关闭资源
//f. 这种机制 非常简洁 程序员不用再为什么时候关闭资源而烦心

