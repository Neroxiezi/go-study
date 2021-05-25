package main

import(
	"fmt"
)

type slice struct {
	Length   int
	Capacity int
	ZerthElement *byte
}

func main() {
	// 基于数组创建一个从 a[strat] 到 a[end -1] 的切片
	a := [5]int{76,77,78,79,80}
	var b []int = a[1:4]
	
	fmt.Println(b)
	// 创建一个长度为4的string数组 并返回一个切片给names
	//names :=[]string{"beijing","上海","PFinal社区"}

	//　创建长度为10 的字符串切片
	str := make([]string,10)
	fmt.Println(str)
}
